//go:generate stringer -type=LogFormat -linecomment
package common

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
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
		return 0, fmt.Errorf("%s is not a valid log format name", str)
	}
}

const (
	DefaultFormat LogFormat = iota // DEFAULT
	TextFormat                     // TEXT
	JsonFormat                     // JSON
	formatEnd
)

func GetLogger(file io.Writer, lv slog.Leveler, format LogFormat, opts ...SlogOption) (logger *slog.Logger, err error) {
	handler, err := GetHandler(file, lv, format, opts...)
	return slog.New(handler), nil
}

func GetHandler(file io.Writer, lv slog.Leveler, format LogFormat, opts ...SlogOption) (handler slog.Handler, err error) {
	options := &slog.HandlerOptions{
		Level:       lv,
		ReplaceAttr: nil,
		AddSource:   false,
	}
	for _, opt := range opts {
		opt.SetOption(options)
	}
	switch format {
	case DefaultFormat:
		handler = NewHandler(file, options)
	case TextFormat:
		handler = slog.NewTextHandler(file, options)
	case JsonFormat:
		handler = slog.NewJSONHandler(file, options)
	default:
		return nil, fmt.Errorf("%d is not a valid log format", format)
	}
	return
}
