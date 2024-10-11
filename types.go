package common

func IsZero[T comparable](a T) bool {
	return a == *new(T)
}
