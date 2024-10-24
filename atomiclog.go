package common

import (
	"context"
	"io"
	"log/slog"
	"sync/atomic"
)

type AtomicHandler struct {
	hdl atomic.Pointer[slog.Handler]
}

func (ah *AtomicHandler) Enabled(ctx context.Context, lv slog.Level) bool {
	return (*(ah.hdl.Load())).Enabled(ctx, lv)
}

func (ah *AtomicHandler) Handle(ctx context.Context, rcd slog.Record) error {
	return (*(ah.hdl.Load())).Handle(ctx, rcd)
}

func (ah *AtomicHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return (*(ah.hdl.Load())).WithAttrs(attrs)

}

func (ah *AtomicHandler) WithGroup(name string) slog.Handler {
	return (*(ah.hdl.Load())).WithGroup(name)
}

func (ah *AtomicHandler) Store(file io.Writer, lv slog.Leveler, format LogFormat, opts ...SlogOption) error {
	handler, err := GetHandler(file, lv, format, opts...)
	if err != nil {
		return err
	}
	ah.hdl.Store(&handler)
	return nil
}
