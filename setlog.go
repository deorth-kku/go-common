package common

import (
	"fmt"
	"io"
	"iter"
	"log/slog"
	"maps"
	"os"
	"strings"
)

var logLevels = map[string]slog.Level{
	"DEBUG":   slog.LevelDebug,
	"ERROR":   slog.LevelError,
	"WARNING": slog.LevelWarn,
	"INFO":    slog.LevelInfo,
}

type EmptyCloser struct{}

func (EmptyCloser) Close() error {
	return nil
}

func SetLog(file string, level string, opts ...SlogOption) (close io.Closer, err error) {
	level = strings.ToUpper(level)
	var f *os.File

	if file == "" {
		f = os.Stderr
		close = EmptyCloser{}
	} else {
		f, err = os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return
		}
		close = f
	}
	options := &slog.HandlerOptions{
		Level:       logLevels[level],
		ReplaceAttr: nil,
		AddSource:   level == "DEBUG",
	}
	for _, opt := range opts {
		opt.SetOption(options)
	}

	var handler slog.Handler = NewHandler(f, options)
	for _, opt := range opts {
		h0 := opt.SetHander(f, options)
		if h0 != nil {
			handler = h0
		}
	}
	slog.SetDefault(slog.New(handler))
	return
}

type SlogOption interface {
	SetOption(opts *slog.HandlerOptions)
	SetHander(w io.Writer, opts *slog.HandlerOptions) slog.Handler
}

type replaceAttr = func(groups []string, attr slog.Attr) slog.Attr

func joinReplaceAttr(a replaceAttr, b replaceAttr) replaceAttr {
	if a == nil && b == nil {
		return nil
	}
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	return func(groups []string, attr slog.Attr) slog.Attr {
		attr = a(groups, attr)
		attr = b(groups, attr)
		return attr
	}
}

type SlogHideTime struct{}

func remove_time(_ []string, attr slog.Attr) slog.Attr {
	if attr.Key == slog.TimeKey {
		return slog.Attr{}
	}
	return attr
}

func (SlogHideTime) SetOption(opts *slog.HandlerOptions) {
	opts.ReplaceAttr = joinReplaceAttr(opts.ReplaceAttr, remove_time)
}

func (SlogHideTime) SetHander(_ io.Writer, _ *slog.HandlerOptions) slog.Handler {
	return nil
}

type SlogMap struct{}

func convert_map(_ []string, attr slog.Attr) slog.Attr {
	m, ok := attr.Value.Any().(map[string]any)
	if ok {
		attr.Value = Map2Group(m)
	}
	return attr
}

func (SlogMap) SetOption(opts *slog.HandlerOptions) {
	opts.ReplaceAttr = joinReplaceAttr(opts.ReplaceAttr, convert_map)
}

func (SlogMap) SetHander(_ io.Writer, _ *slog.HandlerOptions) slog.Handler {
	return nil
}

type SlogIter struct{}

func convert_iter(_ []string, attr slog.Attr) slog.Attr {
	switch it := attr.Value.Any().(type) {
	case iter.Seq2[string, any]:
		attr.Value = Iter2Group(it)
	case func(func(string, any) bool):
		attr.Value = Iter2Group(it)
	}
	return attr
}

func (SlogIter) SetOption(opts *slog.HandlerOptions) {
	opts.ReplaceAttr = joinReplaceAttr(opts.ReplaceAttr, convert_iter)
}

func (SlogIter) SetHander(_ io.Writer, _ *slog.HandlerOptions) slog.Handler {
	return nil
}

type SlogStruct[T any] struct{}

func (SlogStruct[T]) convert_struct(_ []string, attr slog.Attr) slog.Attr {
	a := attr.Value.Any()
	if _, ok := a.(T); ok {
		attr.Value = Map2Group(MustStruct2Map(a))
	}
	return attr
}

func (s SlogStruct[T]) SetOption(opts *slog.HandlerOptions) {
	opts.ReplaceAttr = joinReplaceAttr(opts.ReplaceAttr, s.convert_struct)
}

func (SlogStruct[T]) SetHander(_ io.Writer, _ *slog.HandlerOptions) slog.Handler {
	return nil
}

type SlogText struct{}

func (SlogText) SetOption(opts *slog.HandlerOptions) {
	return
}

func (SlogText) SetHander(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
	return slog.NewTextHandler(w, opts)
}

type SlogJson struct{}

func (SlogJson) SetOption(opts *slog.HandlerOptions) {
	return
}

func (SlogJson) SetHander(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
	return slog.NewJSONHandler(w, opts)
}

var (
	_ SlogOption = SlogText{}
	_ SlogOption = SlogJson{}
	_ SlogOption = SlogHideTime{}
	_ SlogOption = SlogIter{}
)

type antsSlogger struct{}

func (antsSlogger) Printf(format string, args ...any) {
	slog.Debug(fmt.Sprintf(format, args...))
}

var AntsSlogger antsSlogger

func Iter2Group(it iter.Seq2[string, any]) slog.Value {
	values := make([]slog.Attr, 0)
	for k, v := range it {
		switch tv := v.(type) {
		case map[string]any:
			values = append(values, slog.Any(k, Map2Group(tv)))
		case iter.Seq2[string, any]:
			values = append(values, slog.Any(k, Iter2Group(tv)))
		default:
			values = append(values, slog.Any(k, v))
		}
	}
	return slog.GroupValue(values...)
}

func Map2Group(m map[string]any) slog.Value {
	return Iter2Group(maps.All(m))
}
