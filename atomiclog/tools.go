package atomiclog

import (
	"io"
	"log/slog"

	"github.com/deorth-kku/go-common"
)

const (
	ErrNilLogger     = common.ErrorString("nil logger")
	ErrNilHandler    = common.ErrorString("nil log handler")
	ErrNotAtomic     = common.ErrorString("handler is not an AtomicHandler")
	ErrNotLumberjack = common.ErrorString("logger is not a lumberjack")
)

func fromhandler(hdl slog.Handler) (*AtomicHandler, error) {
	if hdl == nil {
		return nil, ErrNilHandler
	}
	ah, ok := hdl.(*AtomicHandler)
	if !ok {
		return nil, ErrNotAtomic
	}
	return ah, nil
}

func fromlogger(logger *slog.Logger) (*AtomicHandler, error) {
	if logger == nil {
		return nil, ErrNilLogger
	}
	return fromhandler(logger.Handler())
}

func Clone(logger *slog.Logger) (*slog.Logger, error) {
	ah, err := fromlogger(logger)
	if err != nil {
		return nil, err
	}
	return slog.New(ah.clone()), nil
}

func Reload[T slog.Handler](logger *slog.Logger, file io.Writer, lv slog.Leveler, format common.LogFormatFunc[T], opts ...common.SlogOption) error {
	ah, err := fromlogger(logger)
	if err != nil {
		return err
	}
	reload(ah, file, lv, format, opts...)
	return nil
}

func SetLevel(logger *slog.Logger, lv slog.Leveler) error {
	ah, err := fromlogger(logger)
	if err != nil {
		return err
	}
	ah.lv.Set(lv.Level())
	return nil
}

type Rotater interface {
	Rotate() error
}

// only works for lumberjack
func Rotate(logger *slog.Logger) error {
	ah, err := fromlogger(logger)
	if err != nil {
		return err
	}
	if lj, ok := ah.f.(Rotater); ok {
		return lj.Rotate()
	}
	return ErrNotLumberjack
}

func GetLogger[T slog.Handler](file io.Writer, lv slog.Leveler, format common.LogFormatFunc[T], opts ...common.SlogOption) *slog.Logger {
	return common.GetLogger(file, lv, NewHandlerFunc(format), opts...)
}

func CloseHandler(hdl slog.Handler) {
	at, err := fromhandler(hdl)
	if err != nil {
		return
	}
	closer, ok := shouldclose(at.f)
	if ok {
		closer.Close()
	}
}
