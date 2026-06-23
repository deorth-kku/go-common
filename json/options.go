package cjson

import (
	"encoding/json/jsontext"
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
