package common

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

var logLevels = map[string]slog.Level{
	"DEBUG":   slog.LevelDebug,
	"ERROR":   slog.LevelError,
	"WARNING": slog.LevelWarn,
	"INFO":    slog.LevelInfo,
}

func SetLog(file string, level string) (err error) {
	level = strings.ToUpper(level)
	var f *os.File

	if file == "" {
		f = os.Stderr
	} else {
		f, err = os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return
		}
	}
	slog.SetDefault(slog.New(slog.NewTextHandler(f, &slog.HandlerOptions{
		Level:       logLevels[level],
		ReplaceAttr: nil,
		AddSource:   level == "DEBUG",
	})))
	return
}

type antsSlogger struct{}

func (antsSlogger) Printf(format string, args ...any) {
	slog.Debug(fmt.Sprintf(format, args...))
}

var AntsSlogger antsSlogger
