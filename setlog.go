package common

import (
	"context"
	"fmt"
	"io"
	"iter"
	"log/slog"
	"maps"
	"os"
	"strconv"
)

func SetLog(file string, level string, format string, opts ...SlogOption) (err error) {
	var f *os.File
	if file == "" {
		f = os.Stderr
	} else {
		f, err = os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return
		}
	}
	var lv slog.Level
	err = lv.UnmarshalText([]byte(level))
	if err != nil {
		return
	}
	fm, err := ParseLogFormat(format)
	if err != nil {
		return
	}
	logger, err := GetLogger(f, lv, fm, opts...)
	if err != nil {
		return
	}
	slog.SetDefault(logger)
	return
}

type SlogOption interface {
	SetOption(opts *slog.HandlerOptions)
}

type replaceAttr = func(groups []string, attr slog.Attr) slog.Attr

func JoinReplaceAttr(a replaceAttr, b replaceAttr) replaceAttr {
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

type SlogAddSource struct{}

func (SlogAddSource) SetOption(opts *slog.HandlerOptions) {
	opts.AddSource = true
}

type SlogAddSourceFunc struct {
	Func func() bool
}

func (s SlogAddSourceFunc) removeSource(groups []string, attr slog.Attr) slog.Attr {
	if len(groups) == 0 && attr.Key == slog.SourceKey && !s.Func() {
		return slog.Attr{}
	}
	return attr
}

func (s SlogAddSourceFunc) SetOption(opts *slog.HandlerOptions) {
	opts.AddSource = true
	opts.ReplaceAttr = JoinReplaceAttr(opts.ReplaceAttr, s.removeSource)
}

type SlogHideTime struct{}

func remove_time(groups []string, attr slog.Attr) slog.Attr {
	if len(groups) == 0 && attr.Key == slog.TimeKey {
		return slog.Attr{}
	}
	return attr
}

func (SlogHideTime) SetOption(opts *slog.HandlerOptions) {
	opts.ReplaceAttr = JoinReplaceAttr(opts.ReplaceAttr, remove_time)
}

type SlogQuoteAttr struct {
	Key    string
	Prefix string
	Suffix string
}

func (s SlogQuoteAttr) quote_attr(groups []string, attr slog.Attr) slog.Attr {
	if len(groups) == 0 && attr.Key == s.Key {
		return slog.String(s.Key, s.Prefix+attr.Value.String()+s.Suffix)
	}
	return attr
}

func (s SlogQuoteAttr) SetOption(opts *slog.HandlerOptions) {
	opts.ReplaceAttr = JoinReplaceAttr(opts.ReplaceAttr, s.quote_attr)
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
	opts.ReplaceAttr = JoinReplaceAttr(opts.ReplaceAttr, convert_map)
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
	opts.ReplaceAttr = JoinReplaceAttr(opts.ReplaceAttr, convert_iter)
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
	opts.ReplaceAttr = JoinReplaceAttr(opts.ReplaceAttr, s.convert_struct)
}

type SlogSlice[T any] struct{}

func sliceconv[Slice ~[]E, E any](s Slice) iter.Seq2[string, any] {
	return func(yield func(string, any) bool) {
		for i, v := range s {
			if !yield(strconv.Itoa(i), v) {
				return
			}
		}
	}
}

func (SlogSlice[T]) convert_slice(_ []string, attr slog.Attr) slog.Attr {
	a := attr.Value.Any()
	if sli, ok := a.([]T); ok {
		attr.Value = Iter2Group(sliceconv(sli))
	}
	return attr
}

func (s SlogSlice[T]) SetOption(opts *slog.HandlerOptions) {
	opts.ReplaceAttr = JoinReplaceAttr(opts.ReplaceAttr, s.convert_slice)
}

type AntsLogger struct {
	*slog.Logger // log to logger
	slog.Leveler // log to level (ants does write with level, you must specify herer)
}

func (al AntsLogger) Printf(format string, args ...any) {
	if al.Logger == nil {
		al.Logger = slog.Default()
	}
	if al.Leveler == nil {
		al.Leveler = slog.LevelInfo
	}
	al.Logger.Log(context.Background(), al.Leveler.Level(), fmt.Sprintf(format, args...))
}

// actually, ants only print log when worker panic, so this is not very useful
var AntsSlogger = AntsLogger{nil, slog.LevelDebug}

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
