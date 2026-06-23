package cjson

import (
	"encoding/json/v2"

	ijson "github.com/deorth-kku/go-common/internal/json"
)

type Unmarshaler[T any] = ijson.Unmarshaler[T]

func UnmarshalV1[T any, P Unmarshaler[T]](data []byte, dst P) error {
	return ijson.UnmarshalV1(data, dst)
}

func MarshalV1[T json.MarshalerTo](v T) ([]byte, error) {
	return ijson.MarshalV1(v)
}
