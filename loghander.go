package common

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"runtime"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"
)

// credit https://groups.google.com/g/golang-nuts/c/aJPXT2NF-Lc/m/bfgqoJSkAQAJ

type MyHandler struct {
	opts      slog.HandlerOptions
	prefix    string // preformatted group names followed by a dot
	preformat string // preformatted Attrs, with an initial space

	mu sync.Mutex
	w  io.Writer
}

func NewHandler(w io.Writer, opts *slog.HandlerOptions) *MyHandler {
	h := &MyHandler{w: w}
	if opts != nil {
		h.opts = *opts
	}
	return h
}

func (h *MyHandler) Enabled(ctx context.Context, level slog.Level) bool {
	minLevel := slog.LevelInfo
	if h.opts.Level != nil {
		minLevel = h.opts.Level.Level()
	}
	return level >= minLevel
}

func (h *MyHandler) WithGroup(name string) slog.Handler {
	return &MyHandler{
		w:         h.w,
		opts:      h.opts,
		preformat: h.preformat,
		prefix:    h.prefix + name + ".",
	}
}

func (h *MyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	var buf []byte
	for _, a := range attrs {
		buf = h.appendAttr(buf, h.prefix, a)
	}
	return &MyHandler{
		w:         h.w,
		opts:      h.opts,
		prefix:    h.prefix,
		preformat: h.preformat + string(buf),
	}
}

func getSource(pc uintptr) string {
	fs := runtime.CallersFrames([]uintptr{pc})
	f, _ := fs.Next()
	return f.File + ":" + strconv.Itoa(f.Line)
}

func (h *MyHandler) Handle(ctx context.Context, r slog.Record) error {
	var buf []byte
	if h.opts.ReplaceAttr != nil {
		key := slog.TimeKey
		val := r.Time.Round(0)
		timeattr := h.opts.ReplaceAttr(nil, slog.Time(key, val))
		if timeattr.Value.Kind() != slog.KindTime || timeattr.Key != slog.TimeKey {
			r.Time = time.Time{}
			h.appendAttr(buf, h.prefix, timeattr)
		} else {
			r.Time = timeattr.Value.Time()
		}
	}
	if !r.Time.IsZero() {
		buf = r.Time.AppendFormat(buf, time.DateTime)
		buf = append(buf, ' ')
	}

	levText := r.Level.String()
	if h.opts.ReplaceAttr != nil {
		lvattr := h.opts.ReplaceAttr(nil, slog.Any(slog.LevelKey, r.Level))
		if lvattr.Key != slog.LevelKey {
			levText = ""
			h.appendAttr(buf, h.prefix, lvattr)
		} else {
			levText = lvattr.Value.String()
		}
	}
	if len(levText) != 0 {
		buf = append(buf, levText...)
		buf = append(buf, ' ')
	}

	if h.opts.AddSource && r.PC != 0 {
		source := getSource(r.PC)
		if h.opts.ReplaceAttr != nil {
			sourceAttr := h.opts.ReplaceAttr(nil, slog.String(slog.SourceKey, source))
			if sourceAttr.Key != slog.SourceKey || sourceAttr.Value.Kind() != slog.KindString {
				source = ""
				h.appendAttr(buf, h.prefix, sourceAttr)
			} else {
				source = sourceAttr.Value.String()
			}
		}
		if len(source) != 0 {
			buf = append(buf, source...)
			buf = append(buf, ' ')
		}
	}

	if h.opts.ReplaceAttr != nil {
		msgattr := h.opts.ReplaceAttr(nil, slog.String(slog.MessageKey, r.Message))
		if msgattr.Value.Kind() != slog.KindString || msgattr.Key != slog.MessageKey {
			r.Message = ""
			h.appendAttr(buf, h.prefix, msgattr)
		} else {
			r.Message = msgattr.Value.String()
		}
	}

	if len(r.Message) != 0 {
		buf = append(buf, r.Message...)
	}

	buf = append(buf, h.preformat...)
	r.Attrs(func(a slog.Attr) bool {
		buf = h.appendAttr(buf, h.prefix, a)
		return true
	})
	buf = append(buf, '\n')
	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := h.w.Write(buf)
	return err
}

func (h *MyHandler) appendAttr(buf []byte, prefix string, a slog.Attr) []byte {
	if h.opts.ReplaceAttr != nil {
		var groups []string
		if prefix != h.prefix {
			groups = strings.Split(prefix[len(h.prefix):], ".")
			groups = slices.DeleteFunc(groups, func(s string) bool {
				return len(s) == 0
			})
		}
		a = h.opts.ReplaceAttr(groups, a)
	}
	if a.Equal(slog.Attr{}) {
		return buf
	}
	switch a.Value.Kind() {
	case slog.KindGroup:
		if a.Key != "" {
			prefix += a.Key + "."
		}
		for _, a := range a.Value.Group() {
			buf = h.appendAttr(buf, prefix, a)
		}
		return buf
	case slog.KindLogValuer:
		a.Value = a.Value.Resolve()
		fallthrough
	default:
		buf = append(buf, ' ')
		buf = append(buf, prefix...)
		buf = append(buf, a.Key...)
		buf = append(buf, '=')
		switch a.Value.Kind() {
		case slog.KindLogValuer:

		}
		return fmt.Appendf(buf, "%v", a.Value.Any())
	}
}
