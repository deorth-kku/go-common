package cslices

import (
	"fmt"
	"testing"

	ctest "github.com/deorth-kku/go-common/test"
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

func TestPartition(t *testing.T) {
	tests := []struct {
		name string
		list []bool
		tr   int
		fl   int
	}{
		{name: "single true", list: []bool{true}, tr: 1, fl: 0},
		{name: "single false", list: []bool{false}, tr: 0, fl: 1},
		{name: "mixed", list: []bool{false, true, false, true}, tr: 2, fl: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := append([]bool(nil), tt.list...)
			tr, fl := PartitionInPlace(list, func(b bool) bool {
				return b
			})
			ctest.Equal(t, len(tr), tt.tr)
			ctest.Equal(t, len(fl), tt.fl)
			for _, v := range tr {
				ctest.True(t, v)
			}
			for _, v := range fl {
				ctest.False(t, v)
			}
		})
	}
}
