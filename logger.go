package common

import (
	"errors"
	"io"
	"log/slog"
	"os"
	"sync/atomic"
)

type Logger struct {
	*slog.Logger
	hdl        *AtomicHandler
	lv         *slog.LevelVar
	f          io.Writer
	clonecount *atomic.Int32
}

func NewLogger(f io.Writer, lv *slog.LevelVar, format LogFormat, opts ...SlogOption) (lg *Logger) {
	lg = new(Logger)
	lg.hdl = new(AtomicHandler)
	lg.Reload(f, lv, format, opts...)
	lg.Logger = slog.New(lg.hdl)
	return
}

func (lg *Logger) Clone() *Logger {
	n := new(Logger)
	lg.clonecount.Add(1)
	n.clonecount = lg.clonecount
	n.f = lg.f
	n.lv = new(slog.LevelVar)
	n.lv.Set(lg.lv.Level())
	n.hdl = lg.hdl.Clone()
	n.Logger = slog.New(n.hdl)
	return n
}

func (lg *Logger) Reload(f io.Writer, lv *slog.LevelVar, format LogFormat, opts ...SlogOption) {
	lg.hdl.Store(f, lv, format, opts...)
	lg.lv = lv
	if closer, ok := lg.f.(io.Closer); ok && lg.f != os.Stderr && lg.f != os.Stdout {
		if lg.clonecount.Load() > 0 {
			lg.clonecount.Add(-1)
		} else {
			closer.Close()
		}
	}
	lg.f = f
	lg.clonecount = new(atomic.Int32)
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
	return errors.New("logger is not a lumberjack")
}
