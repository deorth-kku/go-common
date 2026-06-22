package cslices

import (
	"fmt"
	"testing"
)

func TestCutSlice(t *testing.T) {
	longslice := make([]int, 65535)
	for i := range longslice {
		longslice[i] = i
	}
	last := -1
	for _, subslice := range Cut(longslice, 100) {
		if subslice[0] != last+1 {
			t.Error("no!")
		}
		last = subslice[len(subslice)-1]
	}
}

func TestRand(t *testing.T) {
	s := []string{
		"a",
		"b",
		"c",
	}

	for i, a := range Random(s) {
		fmt.Println(i, a)
	}
	fmt.Println(s)
	Shuffle(s)
	fmt.Println(s)
}
