package atomiclog

import (
	"context"
	"io"
	"log/slog"
	"os"
	"sync/atomic"

	"github.com/deorth-kku/go-common"
)

type AtomicHandler struct {
	hdl    atomic.Pointer[slog.Handler]
	clones *atomic.Int32
	lv     slog.LevelVar
	f      io.Writer
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

func reload[T slog.Handler](ah *AtomicHandler, file io.Writer, lv slog.Leveler, format common.LogFormatFunc[T], opts ...common.SlogOption) {
	if closer, ok := ah.f.(io.Closer); ok && ah.f != os.Stderr && ah.f != os.Stdout {
		if ah.clones.Load() > 0 {
			ah.clones.Add(-1)
		} else {
			defer closer.Close()
		}
	}
	ah.f = file
	ah.lv.Set(lv.Level())
	ah.clones = new(atomic.Int32)
	handler := common.GetHandler(file, lv, format, opts...)
	ah.hdl.Store(&handler)
}

func (ah *AtomicHandler) clone() *AtomicHandler {
	ah.clones.Add(1)
	n := &AtomicHandler{f: ah.f, clones: ah.clones}
	n.hdl.Store(ah.hdl.Load())
	n.lv.Set(ah.lv.Level())
	return n
}

func NewHandlerFunc[T slog.Handler](format common.LogFormatFunc[T]) common.LogFormatFunc[*AtomicHandler] {
	return func(w io.Writer, opts *slog.HandlerOptions) *AtomicHandler {
		ah := &AtomicHandler{
			clones: new(atomic.Int32),
			f:      w,
		}
		if opts == nil {
			opts = new(slog.HandlerOptions)
		}
		if opts.Level != nil {
			ah.lv.Set(opts.Level.Level())
		}
		opts.Level = &ah.lv
		hdl := slog.Handler(format(w, opts))
		ah.hdl.Store(&hdl)
		return ah
	}
}
