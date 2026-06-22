package cslices

import (
	"math/rand/v2"
	"slices"

	citer "github.com/deorth-kku/go-common/iter"
	cmath "github.com/deorth-kku/go-common/math"
)

func Assert[T any](input []any) (output []T) {
	output = make([]T, len(input))
	for i, v := range input {
		output[i] = v.(T)
	}
	return
}

func AssertIter[T any](input []any) citer.Seq[T] {
	return func(yield citer.Yield[T]) {
		for _, i := range input {
			if !yield(i.(T)) {
				return
			}
		}
	}
}

func ToAny[T any, S ~[]T](in S) (out []any) {
	out = make([]any, len(in))
	for i, v := range in {
		out[i] = v
	}
	return
}

func AnyIter[T any, S ~[]T](in S) citer.Seq[any] {
	return func(yield citer.Yield[any]) {
		for _, i := range in {
			if !yield(i) {
				return
			}
		}
	}
}

func Cut[T any, S ~[]T](in S, l int) []S {
	return Collect(slices.Chunk(in, l), cmath.DevidedCeil(len(in), l))
}

// SliceRandom return a iterator of given slice with random order without shuffling the slice
func Random[T any, S ~[]T](in S) citer.Seq2[int, T] {
	idxs := rand.Perm(len(in))
	return func(yield citer.Yield2[int, T]) {
		for _, i := range idxs {
			if !yield(i, in[i]) {
				return
			}
		}
	}
}

func RandElem[T any, S ~[]T](in S) T {
	return in[rand.IntN(len(in))]
}

func Shuffle[T any, S ~[]T](in S) {
	rand.Shuffle(len(in), func(i, j int) {
		in[i], in[j] = in[j], in[i]
	})
}

func Collect[T any](it citer.Seq[T], hint int) (s []T) {
	s = make([]T, 0, hint)
	for i := range it {
		s = append(s, i)
	}
	return
}

func Delete[T comparable, S ~[]T](in S, within ...T) S {
	return slices.DeleteFunc(in, func(v T) bool {
		return slices.Contains(within, v)
	})
}
