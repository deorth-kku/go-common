package common

import (
	"iter"
)

func AnySliceReplaceNil(in []string) (out []any) {
	out = make([]any, len(in))
	for i, arg := range in {
		if len(arg) == 0 {
			out[i] = nil
		} else {
			out[i] = arg
		}
	}
	return
}

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

func SliceAny[T any](in []T) (out []any) {
	out = make([]any, len(in))
	for i, v := range in {
		out[i] = v
	}
	return
}

func SliceAnyIter[T any](in []T) iter.Seq[any] {
	return func(yield func(any) bool) {
		for _, i := range in {
			if !yield(i) {
				return
			}
		}
	}
}

func AnySlice[T any](in []T) []any {
	return SliceAny(in)
}

func CutSlice[t any](in []t, l int) (out [][]t) {
	rounds := len(in) / l
	out = make([][]t, rounds+1)

	for i := 0; i < rounds; i++ {
		out[i] = in[i*l : (i+1)*l]
	}
	out[rounds] = in[rounds*l:]
	return
}
