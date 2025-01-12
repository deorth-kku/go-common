package common

import (
	"iter"
	"math/rand/v2"
	"slices"
)

func SliceAssert[T any](input []any) (output []T) {
	output = make([]T, len(input))
	for i, v := range input {
		output[i] = v.(T)
	}
	return
}

func SliceAssertIter[T any](input []any) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, i := range input {
			if !yield(i.(T)) {
				return
			}
		}
	}
}

func SliceAny[T any, S ~[]T](in S) (out []any) {
	out = make([]any, len(in))
	for i, v := range in {
		out[i] = v
	}
	return
}

func SliceAnyIter[T any, S ~[]T](in S) iter.Seq[any] {
	return func(yield func(any) bool) {
		for _, i := range in {
			if !yield(i) {
				return
			}
		}
	}
}

func AnySlice[T any, S ~[]T](in S) []any {
	return SliceAny(in)
}

func CutSlice[T any, S ~[]T](in S, l int) []S {
	return SliceCollect(slices.Chunk(in, l), DevidedCeil(len(in), l))
}

// SliceRandom return a iterator of given slice with random order without shuffling the slice
func SliceRandom[T any, S ~[]T](in S) iter.Seq2[int, T] {
	idxs := rand.Perm(len(in))
	return func(yield func(int, T) bool) {
		for _, i := range idxs {
			if !yield(i, in[i]) {
				return
			}
		}
	}
}

func SliceShuffle[T any, S ~[]T](in S) {
	rand.Shuffle(len(in), func(i, j int) {
		in[i], in[j] = in[j], in[i]
	})
}

func SliceCollect[T any](it iter.Seq[T], hint int) (s []T) {
	s = make([]T, 0, hint)
	for i := range it {
		s = append(s, i)
	}
	return
}
