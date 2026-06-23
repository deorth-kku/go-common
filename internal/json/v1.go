package ijson

import (
	v1 "encoding/json"
	"encoding/json/v2"
)

type Unmarshaler[T any] interface {
	*T
	json.UnmarshalerFrom
}

func UnmarshalV1[T any, P Unmarshaler[T]](data []byte, dst P) error {
	return json.Unmarshal(data, dst, v1.DefaultOptionsV1())
}

func MarshalV1[T json.MarshalerTo](v T) ([]byte, error) {
	return json.Marshal(v, v1.DefaultOptionsV1())
}
