package common

import (
	"database/sql/driver"
	"encoding"
	"encoding/json"
	"fmt"
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

func (f JsonFloat32[T]) IsZero() bool {
	var t T
	return f == JsonFloat32[T](t.NullValue())
}

func (f *JsonFloat32[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		var t T
		*f = JsonFloat32[T](t.NullValue())
		return nil
	} else {
		return json.Unmarshal(data, (*float32)(f))
	}
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

func (f JsonFloat32[T]) Value() (driver.Value, error) {
	switch {
	case IsNaN(f),
		IsInf(f, 1),
		IsInf(f, -1):
		return nil, nil
	}
	return float32(f), nil
}

func (f *JsonFloat32[T]) Scan(value any) error {
	switch v := value.(type) {
	case nil:
		var t T
		*f = JsonFloat32[T](t.NullValue())
	case float32:
		*f = JsonFloat32[T](v)
	case float64:
		*f = JsonFloat32[T](v)
	default:
		return fmt.Errorf("unsupported value: %v", value)
	}
	return nil
}

type JsonFloat64[T NullTo] float64

func (f JsonFloat64[T]) IsZero() bool {
	var t T
	return f == JsonFloat64[T](t.NullValue())
}

func (f *JsonFloat64[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		var t T
		*f = JsonFloat64[T](t.NullValue())
		return nil
	} else {
		return json.Unmarshal(data, (*float64)(f))
	}
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

func (f JsonFloat64[T]) Value() (driver.Value, error) {
	switch {
	case IsNaN(f),
		IsInf(f, 1),
		IsInf(f, -1):
		return nil, nil
	}
	return float64(f), nil
}

func (f *JsonFloat64[T]) Scan(value any) error {
	switch v := value.(type) {
	case nil:
		var t T
		*f = JsonFloat64[T](t.NullValue())
	case float32:
		*f = JsonFloat64[T](v)
	case float64:
		*f = JsonFloat64[T](v)
	default:
		return fmt.Errorf("unsupported value: %v", value)
	}
	return nil
}

type Nullable[T any] struct {
	V     T
	Valid bool
}

func (nt Nullable[T]) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nt.V)
}

func (nt *Nullable[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		nt.Valid = false
		return nil
	}
	nt.Valid = true
	return json.Unmarshal(data, &nt.V)
}

func (nt Nullable[T]) Value() (driver.Value, error) {
	if nt.Valid {
		return nt.V, nil
	}
	return nil, nil
}

func (nt *Nullable[T]) Scan(value any) error {
	switch v := value.(type) {
	case nil:
		return nil
	case T:
		nt.V = v
		nt.Valid = true
	default:
		return fmt.Errorf("unsupported value: %v", value)
	}
	return nil
}

func NewNullable[T any](v T) Nullable[T] {
	return Nullable[T]{V: v, Valid: true}
}

type SqlString[T encoding.TextMarshaler] struct {
	Raw T
}

func (f SqlString[T]) Value() (driver.Value, error) {
	return f.Raw.MarshalText()
}

func unmarshalText[T any](f *T, text []byte) error {
	unmalshaler, ok := any(f).(encoding.TextUnmarshaler)
	if !ok {
		return fmt.Errorf("unsupported value type: %T", *f)
	}
	return unmalshaler.UnmarshalText(text)
}

func (f *SqlString[T]) Scan(value any) error {
	switch v := value.(type) {
	case []byte:
		return unmarshalText(&f.Raw, v)
	case string:
		return unmarshalText(&f.Raw, []byte(v))
	default:
		return fmt.Errorf("unsupported value: %v", value)
	}
}

type SqlNullString[T encoding.TextMarshaler] struct {
	Raw   T
	Valid bool
}

func (f SqlNullString[T]) Value() (driver.Value, error) {
	if !f.Valid {
		return nil, nil
	}
	return f.Raw.MarshalText()
}

func (f *SqlNullString[T]) Scan(value any) error {
	switch v := value.(type) {
	case nil:
		f.Valid = false
		return nil
	case []byte:
		f.Valid = true
		return unmarshalText(&f.Raw, v)
	case string:
		f.Valid = true
		return unmarshalText(&f.Raw, []byte(v))
	default:
		return fmt.Errorf("unsupported value: %v", value)
	}
}
