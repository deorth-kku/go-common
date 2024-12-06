package common

import (
	"math"
	"math/rand/v2"
	"reflect"
	"strconv"
	"unsafe"
)

type SignedInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type UnsignedInt interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type AnyInt interface {
	SignedInt | UnsignedInt
}

type Number interface {
	AnyInt | Float
}

const (
	uvnan32 = 0x7FC00001
	uvnan64 = 0x7FF8000000000001
)

func Parse[T Number](s string, base int) (t T, err error) {
	size := int(unsafe.Sizeof(t))
	switch size {
	case 1:
		p := (*uint8)(unsafe.Pointer(&t))
		*p = math.MaxUint8
	case 2:
		p := (*uint16)(unsafe.Pointer(&t))
		*p = math.MaxUint16
	case 4:
		p := (*uint32)(unsafe.Pointer(&t))
		*p = uvnan32
		if t != t {
			goto float
		}
		*p = math.MaxUint32
	case 8:
		p := (*uint64)(unsafe.Pointer(&t))
		*p = uvnan64
		if t != t {
			goto float
		}
		*p = math.MaxUint64
	default:
		return 0, ErrorString("unexpect type " + reflect.TypeOf(t).Name())
	}
	if t < 0 {
		var i64 int64
		i64, err = strconv.ParseInt(s, base, size*8)
		return T(i64), err
	} else {
		var u64 uint64
		u64, err = strconv.ParseUint(s, base, size*8)
		return T(u64), err
	}
float:
	var f64 float64
	f64, err = strconv.ParseFloat(s, size*8)
	return T(f64), err
}

func DevidedCeil[T AnyInt](a, b T) T {
	return T(math.Ceil(float64(a) / float64(b)))
}

func MaxInt[T AnyInt]() (t T) {
	size := unsafe.Sizeof(t) * 8
	if ^t < 0 {
		return T(1)<<(size-1) - 1 // signed
	}
	return ^t // unsigned
}

func Roll(n int) int {
	if n < 1 {
		return n
	}
	return rand.IntN(n)
}

func Nan32() float32 {
	return math.Float32frombits(uvnan32)
}

func IsNaN[F Float](f F) bool {
	return f != f
}
