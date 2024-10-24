package common

import (
	"io"
	"log/slog"
	"os"
)

type Logger struct {
	*slog.Logger
	hdl *AtomicHandler
	lv  *slog.LevelVar
	f   io.WriteCloser
}

func NewLogger(f io.WriteCloser, lv *slog.LevelVar, format LogFormat, opts ...SlogOption) (lg *Logger) {
	lg = new(Logger)
	lg.hdl = new(AtomicHandler)
	lg.Reload(f, lv, format, opts...)
	lg.Logger = slog.New(lg.hdl)
	return
}

func (lg *Logger) Reload(f io.WriteCloser, lv *slog.LevelVar, format LogFormat, opts ...SlogOption) {
	lg.hdl.Store(f, lv, format, opts...)
	lg.lv = lv
	if lg.f != nil && lg.f != os.Stderr && lg.f != os.Stdout {
		lg.f.Close()
	}
	lg.f = f
	return
}

func (lg *Logger) SetLevel(l slog.Leveler) {
	lg.lv.Set(l.Level())
}

type Rotater interface {
	Rotate() error
}

// only works for lumberjack
func (lg *Logger) Rotate() error {
	if lj, ok := lg.f.(Rotater); ok {
		return lj.Rotate()
	}
	return nil
}
