package common

import "iter"

type Pair[K any, V any] struct {
	Key   K
	Value V
}

func NewPair[K any, V any](Key K, Value V) Pair[K, V] {
	return Pair[K, V]{Key: Key, Value: Value}
}

type PairSlice[K any, V any] []Pair[K, V]

func (ps PairSlice[K, V]) Range(yield func(K, V) bool) {
	for _, pair := range ps {
		if !yield(pair.Key, pair.Value) {
			return
		}
	}
}

func EmptyRange[T any](func(T) bool)            {}
func EmptyRange2[K any, V any](func(K, V) bool) {}

func SafeRange[T any](f iter.Seq[T]) iter.Seq[T] {
	if f == nil {
		return EmptyRange
	}
	return f
}

func SafeRange2[K any, V any](f iter.Seq2[K, V]) iter.Seq2[K, V] {
	if f == nil {
		return EmptyRange2
	}
	return f
}
