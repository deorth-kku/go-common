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
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type AnyInt interface {
	SignedInt | UnsignedInt
}

type Number interface {
	AnyInt | Float
}

func Abs[T SignedInt | Float](i T) T {
	if i >= 0 {
		return i
	}
	return -i
}

const (
	typeFloat = iota
	typeSigned
	typeUnsigned
)

func getTinfo[T Number]() (int, int) {
	var t T
	size := unsafe.Sizeof(t)
	switch size {
	case 1:
		*(*uint8)(unsafe.Pointer(&t)) = math.MaxUint8
	case 2:
		*(*uint16)(unsafe.Pointer(&t)) = math.MaxUint16
	case 4:
		*(*uint32)(unsafe.Pointer(&t)) = math.MaxUint32
	case 8:
		*(*uint64)(unsafe.Pointer(&t)) = math.MaxUint64
	default:
		panic("unexpect type " + reflect.TypeOf(t).Name())
	}
	if t != t {
		return typeFloat, int(size * 8)
	}
	if t < 0 {
		return typeSigned, int(size * 8)
	}
	return typeUnsigned, int(size * 8)
}

func Parse[T Number](s string, base int) (T, error) {
	t, size := getTinfo[T]()
	switch t {
	case typeFloat:
		f64, err := strconv.ParseFloat(s, size)
		return T(f64), err
	case typeSigned:
		i64, err := strconv.ParseInt(s, base, size)
		return T(i64), err
	case typeUnsigned:
		u64, err := strconv.ParseUint(s, base, size)
		return T(u64), err
	default:
		panic("unexpect type " + reflect.TypeFor[T]().Name())
	}
}

func Format[T Number](n T, base int) string {
	t, size := getTinfo[T]()
	switch t {
	case typeFloat:
		return strconv.FormatFloat(float64(n), 'f', -1, size)
	case typeSigned:
		return strconv.FormatInt(int64(n), 10)
	case typeUnsigned:
		return strconv.FormatUint(uint64(n), 10)
	default:
		panic("unexpect type " + reflect.TypeFor[T]().Name())
	}
}

func DevidedCeil[T AnyInt](a, b T) T {
	if (a < 0) == (b < 0) && a%b != 0 {
		return a/b + 1
	}
	return a / b
}

func MaxInt[T AnyInt]() (t T) {
	size := unsafe.Sizeof(t) * 8
	if ^t < 0 {
		return T(1)<<(size-1) - 1 // signed
	}
	return ^t // unsigned
}

func Roll[T AnyInt](n T) T {
	if n < 1 {
		return n
	}
	return rand.N(n)
}

const (
	uvnan32    = 0x7FC00001
	uninf32    = 0x7F800000
	uvneginf32 = 0xFF800000
)

func NaN32() float32 {
	return math.Float32frombits(uvnan32)
}

func NaN[T Float]() T {
	switch unsafe.Sizeof(T(0)) {
	case 4:
		return T(NaN32())
	case 8:
		return T(math.NaN())
	default:
		panic("unexpected float size " + strconv.Itoa(int(unsafe.Sizeof(T(0)))))
	}
}

func IsNaN[F Float](f F) bool {
	return f != f
}

func Inf32(sign int) float32 {
	var v uint32
	if sign >= 0 {
		v = uninf32
	} else {
		v = uvneginf32
	}
	return math.Float32frombits(v)
}

func Inf[T Float](sign int) T {
	switch unsafe.Sizeof(T(0)) {
	case 4:
		return T(Inf32(sign))
	case 8:
		return T(math.Inf(sign))
	default:
		panic("unexpected float size " + strconv.Itoa(int(unsafe.Sizeof(T(0)))))
	}
}

func IsInf[F Float](f F, sign int) bool {
	switch unsafe.Sizeof(f) {
	case 4:
		return sign >= 0 && f > math.MaxFloat32 || sign <= 0 && f < -math.MaxFloat32
	case 8:
		return math.IsInf(float64(f), sign)
	default:
		panic("unexpected float size " + strconv.Itoa(int(unsafe.Sizeof(F(0)))))
	}
}

func Sign[T Number](i T) int {
	switch {
	case i == 0:
		return 0
	case i > 0:
		return 1
	default:
		return -1
	}
}

func BoolToNumber[T Number](b bool) T {
	if b {
		return 1
	}
	return 0
}
