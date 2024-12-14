package common

import "iter"

type Pair[K any, V any] struct {
	Key   K
	Value V
}

func NewPair[K any, V any](Key K, Value V) Pair[K, V] {
	return Pair[K, V]{Key: Key, Value: Value}
}

func PairSliceCollect[K any, V any](iter iter.Seq2[K, V]) (pairs PairSlice[K, V]) {
	for k, v := range iter {
		pairs = append(pairs, NewPair(k, v))
	}
	return
}
func PairSliceFromMap[K comparable, V any, M ~map[K]V](m M) (pairs PairSlice[K, V]) {
	pairs = make(PairSlice[K, V], 0, len(m))
	for k, v := range m {
		pairs = append(pairs, NewPair(k, v))
	}
	return
}

type PairSlice[K any, V any] []Pair[K, V]

func (ps PairSlice[K, V]) Keys(yield func(K) bool) {
	for _, pair := range ps {
		if !yield(pair.Key) {
			return
		}
	}
}

func (ps PairSlice[K, V]) Values(yield func(V) bool) {
	for _, pair := range ps {
		if !yield(pair.Value) {
			return
		}
	}
}

func (ps PairSlice[K, V]) Range(yield func(K, V) bool) {
	for _, pair := range ps {
		if !yield(pair.Key, pair.Value) {
			return
		}
	}
}

func EmptyRange[T any](func(T) bool)            {}
func EmptyRange2[K any, V any](func(K, V) bool) {}

func SafeRange[T any](fs ...iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, f := range fs {
			if f == nil {
				continue
			}
			for v := range f {
				if !yield(v) {
					return
				}
			}
		}
	}
}

func SafeRange2[K any, V any](fs ...iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, f := range fs {
			if f == nil {
				continue
			}
			for k, v := range f {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}
