package common

type Pair[K any, V any] struct {
	Key   K
	Value V
}

func NewPair[K any, V any](Key K, Value V) Pair[K, V] {
	return Pair[K, V]{Key: Key, Value: Value}
}

type PairSlice[K any, V any] []Pair[K, V]

func (ps PairSlice[K, V]) Range(yield Yield2[K, V]) {
	for _, pair := range ps {
		if !yield(pair.Key, pair.Value) {
			return
		}
	}
}

func EmptyRange[T any](Yield[T])             {}
func EmptyRange2[K any, V any](Yield2[K, V]) {}

func SafeRange[T any](f Seq[T]) Seq[T] {
	if f == nil {
		return EmptyRange
	}
	return f
}

func SafeRange2[K any, V any](f Seq2[K, V]) Seq2[K, V] {
	if f == nil {
		return EmptyRange2
	}
	return f
}

type (
	Yield[T any]         = func(T) bool
	Yield2[K any, V any] = func(K, V) bool
	Seq[T any]           = func(Yield[T])
	Seq2[K any, V any]   = func(Yield2[K, V])
)
