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
