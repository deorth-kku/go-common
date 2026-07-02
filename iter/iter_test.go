package citer

import (
	"fmt"
	"maps"
	"slices"
	"testing"
)

func TestSeq2(t *testing.T) {
	keys := slices.Collect(Seq2K(maps.All(map[string]int{
		"1": 2,
		"3": 4,
	})))
	fmt.Println(keys)
}

func TestIndex(t *testing.T) {
	keys := maps.Collect(Count[int](slices.Values([]string{"0", "1", "2", "3", "4"})))
	fmt.Println(keys)
}
