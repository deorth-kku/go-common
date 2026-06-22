package citer

import (
	"maps"
	"slices"
	"testing"
)

func TestSeq2(t *testing.T) {
	keys := slices.Collect(Seq2K(maps.All(map[string]int{
		"1": 2,
		"3": 4,
	})))
	println(keys)
}
