package common

import (
	"math"
	"math/rand"
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
	return rand.Intn(n)
}
