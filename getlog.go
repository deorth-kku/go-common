package common

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"log/slog"
)

type LogFormat uint8

func (lf *LogFormat) UnmarshalJSON(data []byte) (err error) {
	var i uint8
	err = json.Unmarshal(data, &i)
	if err == nil {
		*lf = LogFormat(i)
		if *lf >= formatEnd {
			return fmt.Errorf("%d is not a valid log format", i)
		}
		return
	}
	var str string
	err = json.Unmarshal(data, &str)
	if err != nil {
		return
	}
	*lf, err = ParseLogFormat(str)
	return
}

func ParseLogFormat(str string) (lf LogFormat, err error) {
	switch strings.ToUpper(str) {
	case "", "DEFAULT":
		lf = DefaultFormat
	case "TEXT":
		lf = TextFormat
	case "JSON":
		lf = JsonFormat
	default:
		return 0, fmt.Errorf("%s is not a valid log format name", str)
	}
	return
}

const (
	DefaultFormat LogFormat = iota
	TextFormat
	JsonFormat
	formatEnd
)

func GetLogger(file io.Writer, lv slog.Leveler, format LogFormat, opts ...SlogOption) (logger *slog.Logger, err error) {
	options := &slog.HandlerOptions{
		Level:       lv,
		ReplaceAttr: nil,
		AddSource:   false,
	}
	for _, opt := range opts {
		opt.SetOption(options)
	}
	var handler slog.Handler
	switch format {
	case DefaultFormat:
		handler = NewHandler(file, options)
	case TextFormat:
		handler = slog.NewTextHandler(file, options)
	case JsonFormat:
		handler = slog.NewJSONHandler(file, options)
	default:
		err = fmt.Errorf("%d is not a valid log format", format)
		return
	}
	return slog.New(handler), nil
}
