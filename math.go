package common

import (
	"math"
	"math/rand/v2"
)

type AnyInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func DevidedCeil[T AnyInt](a, b T) T {
	return T(math.Ceil(float64(a) / float64(b)))
}

func Roll(n int) int {
	if n < 1 {
		return n
	}
	return rand.IntN(n)
}

func Nan32() float32 {
	return float32(math.NaN())
}

func IsNaN[F float](f F) bool {
	return f != f
}
