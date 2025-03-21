package common

func IsZero[T comparable](a T) bool {
	return a == *new(T)
}

func Equal[T comparable](a, b T) bool {
	return a == b
}

type CanEqual[T any] interface {
	Equal(T) bool
}

func EqualT[T CanEqual[T]](a, b T) bool {
	return a.Equal(b)
}

func IsZeroT[T CanEqual[T]](a T) bool {
	return a.Equal(*new(T))
}

type CanCompare[T any] interface {
	Compare(T) int
}

func CompareT[T CanCompare[T]](a, b T) int {
	return a.Compare(b)
}
