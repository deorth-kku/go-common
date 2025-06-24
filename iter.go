package common

import (
	"maps"
	"slices"
)

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

type Pair[K any, V any] struct {
	Key   K
	Value V
}

func NewPair[K any, V any](Key K, Value V) Pair[K, V] {
	return Pair[K, V]{Key: Key, Value: Value}
}

func PairSliceCollect[K any, V any](iter Seq2[K, V], hint ...int) (pairs PairSlice[K, V]) {
	if len(hint) != 0 {
		pairs = make(PairSlice[K, V], 0, hint[0])
	}
	for k, v := range iter {
		pairs = append(pairs, NewPair(k, v))
	}
	return
}
func PairSliceFromMap[K comparable, V any, M ~map[K]V](m M) (pairs PairSlice[K, V]) {
	return PairSliceCollect(maps.All(m), len(m))
}

type PairSlice[K any, V any] []Pair[K, V]

func (ps PairSlice[K, V]) Keys(yield Yield[K]) {
	slices.Values(ps)(func(line Pair[K, V]) bool {
		return yield(line.Key)
	})
}

func (ps PairSlice[K, V]) SliceKeys() []K {
	return SliceCollect(ps.Keys, len(ps))
}

func (ps PairSlice[K, V]) Values(yield Yield[V]) {
	slices.Values(ps)(func(line Pair[K, V]) bool {
		return yield(line.Value)
	})
}

func (ps PairSlice[K, V]) SliceValues() []V {
	return SliceCollect(ps.Values, len(ps))
}

func (ps PairSlice[K, V]) Range(yield Yield2[K, V]) {
	slices.Values(ps)(func(line Pair[K, V]) bool {
		return yield(line.Key, line.Value)
	})
}

func (ps PairSlice[K, V]) Backward(yield Yield2[K, V]) {
	slices.Backward(ps)(func(_ int, line Pair[K, V]) bool {
		return yield(line.Key, line.Value)
	})
}

func (ps PairSlice[K, V]) Search(key K, equal func(a, b K) bool) (V, bool) {
	for _, line := range ps {
		if equal(line.Key, key) {
			return line.Value, true
		}
	}
	var v V
	return v, false
}

func Search[K comparable, V any](ps PairSlice[K, V], key K) (V, bool) {
	return ps.Search(key, Equal)
}

func SearchEqual[K CanEqual[K], V any](ps PairSlice[K, V], key K) (V, bool) {
	return ps.Search(key, EqualT)
}

func EmptyRange[T any](Yield[T])             {}
func EmptyRange2[K any, V any](Yield2[K, V]) {}

func SafeRange[T any](fs ...Seq[T]) Seq[T] {
	return func(yield Yield[T]) {
		for i, f := range fs {
			if f == nil {
				continue
			}
			if i == len(fs)-1 {
				f(yield)
				return
			}
			for v := range f {
				if !yield(v) {
					return
				}
			}
		}
	}
}

func SafeRange2[K any, V any](fs ...Seq2[K, V]) Seq2[K, V] {
	return func(yield Yield2[K, V]) {
		for i, f := range fs {
			if f == nil {
				continue
			}
			if i == len(fs)-1 {
				f(yield)
				return
			}
			for k, v := range f {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}

type PairChan[K any, V any] chan Pair[K, V]

func (pc PairChan[K, V]) Range(yield Yield2[K, V]) {
	for pair := range pc {
		if !yield(pair.Key, pair.Value) {
			return
		}
	}
}

func Seq2K[K any, V any](it Seq2[K, V]) Seq[K] {
	return func(yield Yield[K]) {
		it(func(k K, _ V) bool {
			return yield(k)
		})
	}
}

func Seq2V[K any, V any](it Seq2[K, V]) Seq[V] {
	return func(yield Yield[V]) {
		it(func(_ K, v V) bool {
			return yield(v)
		})
	}
}

func Filter[T any](seq Seq[T], filter func(T) bool) Seq[T] {
	return func(yield Yield[T]) {
		seq(func(v T) bool {
			if filter(v) {
				return yield(v)
			}
			return true
		})
	}
}
func Filter2[K any, V any](seq Seq2[K, V], filter func(K, V) bool) Seq2[K, V] {
	return func(yield Yield2[K, V]) {
		seq(func(k K, v V) bool {
			if filter(k, v) {
				return yield(k, v)
			}
			return true
		})
	}
}
