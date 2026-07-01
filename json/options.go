package cjson

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"reflect"

	"github.com/deorth-kku/go-common/args"
)

const omitTopLevelNewline = 1 << 4

var jsonflagsBoolType = reflect.TypeOf(jsontext.AllowDuplicateNames(true))

func WithOmitTopLevelNewline(b bool) jsontext.Options {
	rv := reflect.New(jsonflagsBoolType).Elem()
	if b {
		rv.SetUint(omitTopLevelNewline | 1)
	} else {
		rv.SetUint(omitTopLevelNewline | 0)
	}
	return args.MustOk(reflect.TypeAssert[jsontext.Options](rv))
}

type (
	UnmarshalerFunc[T any] = func(*jsontext.Decoder, T) error
	MarshalerFunc[T any]   = func(*jsontext.Encoder, T) error
)

func UpdateUnmarshalers[T any](opts json.Options, fn UnmarshalerFunc[T]) json.Options {
	unmarshalers, ok := json.GetOption(opts, json.WithUnmarshalers)
	if ok {
		unmarshalers = json.JoinUnmarshalers(unmarshalers, json.UnmarshalFromFunc(fn))
	} else {
		unmarshalers = json.UnmarshalFromFunc(fn)
	}
	return json.JoinOptions(opts, json.WithUnmarshalers(unmarshalers))
}

func UpdateMarshalers[T any](opts json.Options, fn MarshalerFunc[T]) json.Options {
	marshalers, ok := json.GetOption(opts, json.WithMarshalers)
	if ok {
		marshalers = json.JoinMarshalers(marshalers, json.MarshalToFunc(fn))
	} else {
		marshalers = json.MarshalToFunc(fn)
	}
	return json.JoinOptions(opts, json.WithMarshalers(marshalers))
}
