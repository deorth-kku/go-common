package common

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"runtime"
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

func (h *MyHandler) Handle(ctx context.Context, r slog.Record) error {
	var buf []byte

	if h.opts.ReplaceAttr != nil {
		key := slog.TimeKey
		val := r.Time.Round(0)
		timeattr := h.opts.ReplaceAttr(nil, slog.Time(key, val))
		var ok bool
		if r.Time, ok = timeattr.Value.Any().(time.Time); !ok {
			h.appendAttr(buf, h.prefix, timeattr)
		}
	}
	if !r.Time.IsZero() {
		buf = r.Time.AppendFormat(buf, time.DateTime)
		buf = append(buf, ' ')
	}

	levText := (r.Level.String() + " ")[0:5]

	buf = append(buf, levText...)
	buf = append(buf, ' ')
	if h.opts.AddSource && r.PC != 0 {
		fs := runtime.CallersFrames([]uintptr{r.PC})
		f, _ := fs.Next()
		buf = append(buf, f.File...)
		buf = append(buf, ':')
		buf = strconv.AppendInt(buf, int64(f.Line), 10)
		buf = append(buf, ' ')
	}
	buf = append(buf, r.Message...)
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
			groups = strings.Split(prefix[len(h.prefix)+1:], ".")
		}
		a = h.opts.ReplaceAttr(groups, a)
	}
	if a.Equal(slog.Attr{}) {
		return buf
	}
	if a.Value.Kind() != slog.KindGroup {
		buf = append(buf, ' ')
		buf = append(buf, prefix...)
		buf = append(buf, a.Key...)
		buf = append(buf, '=')
		return fmt.Appendf(buf, "%v", a.Value.Any())
	}
	// Group
	if a.Key != "" {
		prefix += a.Key + "."
	}
	for _, a := range a.Value.Group() {
		buf = h.appendAttr(buf, prefix, a)
	}
	return buf
}
