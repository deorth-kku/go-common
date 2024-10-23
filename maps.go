package common

import (
	"maps"
	"slices"
)

func MapKeys[K comparable, V any](in map[K]V) (out []K) {
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

func MapKeysSort[K comparable, V any](in map[K]V, sort_func func(a, b K) int) (out []K) {
	out = MapKeys(in)
	slices.SortFunc(out, sort_func)
	return
}

// MapMerge return a new map with all key-value pairs from the sources maps, the new keys and values are set using ordinary assignment.
func MapMerge[K comparable, V any](in ...map[K]V) (out map[K]V) {
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

func MapAssert[K comparable, V any](input map[K]any) (output map[K]V) {
	output = make(map[K]V, len(input))
	for k, v := range input {
		output[k] = v.(V)
	}
	return
}

func MapAssertIter[K comparable, V any](input map[K]any) Seq2[K, V] {
	return func(yield Yield2[K, V]) {
		for k, v := range input {
			if !yield(k, v.(V)) {
				return
			}
		}
	}
}

func MapAny[K comparable, T any](input map[K]T) (output map[K]any) {
	output = make(map[K]any, len(input))
	for k, v := range input {
		output[k] = v
	}
	return
}

func MapAnyIter[K comparable, T any](input map[K]T) Seq2[K, any] {
	return func(yield Yield2[K, any]) {
		for k, v := range input {
			if !yield(k, v) {
				return
			}
		}
	}
}
