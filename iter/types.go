package citer

type (
	Yield[T any]     = func(T) bool
	Yield2[K, V any] = func(K, V) bool
	Seq[T any]       = func(Yield[T])
	Seq2[K, V any]   = func(Yield2[K, V])
	Ranger[T any]    interface {
		Range(Yield[T])
	}
	Ranger2[K, V any] = interface {
		Range(Yield2[K, V])
	}
)
