package citer

import (
	cmath "github.com/deorth-kku/go-common/math"
)

func EmptyRange[T any](Yield[T])             {}
func EmptyRange2[K any, V any](Yield2[K, V]) {}

func SafeRange[IT ~Seq[T], T any](fs ...IT) IT {
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

func SafeRange2[IT ~Seq2[K, V], K, V any](fs ...IT) IT {
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

func Filter[T any](seq Seq[T], filter Yield[T]) Seq[T] {
	return func(yield Yield[T]) {
		seq(func(v T) bool {
			if filter(v) {
				return yield(v)
			}
			return true
		})
	}
}
func Filter2[K any, V any](seq Seq2[K, V], filter Yield2[K, V]) Seq2[K, V] {
	return func(yield Yield2[K, V]) {
		seq(func(k K, v V) bool {
			if filter(k, v) {
				return yield(k, v)
			}
			return true
		})
	}
}

func Count[I cmath.AnyInt, T any](seq Seq[T]) Seq2[I, T] {
	return func(y Yield2[I, T]) {
		var i I
		seq(func(v T) bool {
			b := y(i, v)
			i++
			return b
		})
	}
}
