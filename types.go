package common

func IsZero[T comparable](a T) bool {
	return a == *new(T)
}

type ErrorString string

func (e ErrorString) Error() string {
	return string(e)
}
