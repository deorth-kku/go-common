package cmaps

import (
	"maps"
	"slices"

	citer "github.com/deorth-kku/go-common/iter"
)

func Keys[K comparable, V any, M ~map[K]V](in M) (out []K) {
	if in == nil {
		return nil
	}
	out = make([]K, len(in))
	i := 0
	for k := range in {
		out[i] = k
		i++
	}

	return
}

func KeysSort[K comparable, V any, M ~map[K]V](in M, sort_func func(a, b K) int) (out []K) {
	out = Keys(in)
	slices.SortFunc(out, sort_func)
	return
}

// MapMerge return a new map with all key-value pairs from the sources maps, the new keys and values are set using ordinary assignment.
func Merge[K comparable, V any, M ~map[K]V](in ...M) (out M) {
	switch len(in) {
	case 0:
		return nil
	case 1:
		return maps.Clone(in[0])
	default:
		out = make(map[K]V)
		for _, m := range in {
			maps.Copy(out, m)
		}
		return
	}
}

func Assert[K comparable, V any, M ~map[K]any](input M) (output map[K]V) {
	output = make(map[K]V, len(input))
	for k, v := range input {
		output[k] = v.(V)
	}
	return
}

func AssertIter[K comparable, V any, M ~map[K]any](input M) citer.Seq2[K, V] {
	return func(yield citer.Yield2[K, V]) {
		for k, v := range input {
			if !yield(k, v.(V)) {
				return
			}
		}
	}
}

func Any[K comparable, T any, M ~map[K]T](input M) (output map[K]any) {
	output = make(map[K]any, len(input))
	for k, v := range input {
		output[k] = v
	}
	return
}

func AnyIter[K comparable, T any, M ~map[K]T](input M) citer.Seq2[K, any] {
	return func(yield citer.Yield2[K, any]) {
		for k, v := range input {
			if !yield(k, v) {
				return
			}
		}
	}
}

func Collect[K comparable, V any](it citer.Seq2[K, V], hint int) (m map[K]V) {
	m = make(map[K]V, hint)
	for k, v := range it {
		m[k] = v
	}
	return
}
