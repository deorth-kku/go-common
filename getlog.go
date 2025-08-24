package common

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"

	"log/slog"
)

//go:generate stringer -type=LogFormat -linecomment  --trimprefix formatEnd
type LogFormat uint8

func (lf *LogFormat) UnmarshalJSON(data []byte) (err error) {
	err = json.Unmarshal(data, (*uint8)(lf))
	if err == nil {
		return lf.Validate()
	}
	var str string
	err = json.Unmarshal(data, &str)
	if err != nil {
		return
	}
	*lf, err = ParseLogFormat(str)
	return
}

func (lf *LogFormat) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(lf.String())), nil
}

func ParseLogFormat(str string) (LogFormat, error) {
	switch strings.ToUpper(str) {
	case "", DefaultFormat.String():
		return DefaultFormat, nil
	case TextFormat.String():
		return TextFormat, nil
	case JsonFormat.String():
		return JsonFormat, nil
	default:
		return 0, fmt.Errorf("%s is not a valid LogFormat name", str)
	}
}

const (
	DefaultFormat LogFormat = iota // DEFAULT
	TextFormat                     // TEXT
	JsonFormat                     // JSON
	formatEnd
)

func (f LogFormat) Validate() error {
	if f >= formatEnd {
		return fmt.Errorf("%d is not a valid LogFormat", f)
	}
	return nil
}

func (f LogFormat) NewHandler(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
	switch f {
	case DefaultFormat:
		return NewHandler(w, opts)
	case TextFormat:
		return slog.NewTextHandler(w, opts)
	case JsonFormat:
		return slog.NewJSONHandler(w, opts)
	default:
		return nil
	}
}

type LogFormatFunc[T slog.Handler] = func(w io.Writer, opts *slog.HandlerOptions) T

func GetHandler[T slog.Handler](file io.Writer, lv slog.Leveler, format LogFormatFunc[T], opts ...SlogOption) slog.Handler {
	options := &slog.HandlerOptions{
		Level:       lv,
		ReplaceAttr: nil,
		AddSource:   false,
	}
	for _, opt := range opts {
		opt.SetOption(options)
	}
	return format(file, options)
}

func GetLogger[T slog.Handler](file io.Writer, lv slog.Leveler, format LogFormatFunc[T], opts ...SlogOption) *slog.Logger {
	return slog.New(GetHandler(file, lv, format, opts...))
}
