package common

import (
	"encoding/json"
	"math"
)

type (
	ToPosInf struct{}
	ToNegInf struct{}
	ToNaN    struct{}
	ToZero   struct{}
)

func (ToPosInf) NullValue() float64 {
	return math.Inf(1)
}

func (ToNegInf) NullValue() float64 {
	return math.Inf(-1)
}

func (ToNaN) NullValue() float64 {
	return math.NaN()
}

func (ToZero) NullValue() float64 {
	return 0
}

type NullTo interface {
	NullValue() float64
}

type JsonFloat32[T NullTo] float32

func (f *JsonFloat32[T]) UnmarshalJSON(data []byte) error {
	var t T
	v := float32(t.NullValue())
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*f = JsonFloat32[T](v)
	return nil
}

func (f JsonFloat32[T]) MarshalJSON() ([]byte, error) {
	switch {
	case IsNaN(f),
		IsInf(f, 1),
		IsInf(f, -1):
		return []byte("null"), nil
	}
	return json.Marshal(float32(f))
}

type JsonFloat64[T NullTo] float64

func (f *JsonFloat64[T]) UnmarshalJSON(data []byte) error {
	var t T
	v := t.NullValue()
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*f = JsonFloat64[T](v)
	return nil
}

func (f JsonFloat64[T]) MarshalJSON() ([]byte, error) {
	switch {
	case IsNaN(f),
		IsInf(f, 1),
		IsInf(f, -1):
		return []byte("null"), nil
	}
	return json.Marshal(float64(f))
}
