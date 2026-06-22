package datatypes

import (
	"cmp"
	"math"
	"math/rand/v2"
	"slices"
	"testing"

	cslices "github.com/deorth-kku/go-common/datatypes/slices"
)

func TestBSMap(t *testing.T) {
	m := NewBSMap[int, Empty]()
	for range 100 {
		m.Store(rand.Int(), Empty{})
	}
	if !slices.IsSorted(m.SliceKeys()) {
		t.Fatal("not sorted")
	}
	for key := range m.Keys {
		_, ok := m.Load(key)
		if !ok {
			t.Fatal("not found")
		}
	}
	for _, key := range cslices.Random(m.SliceKeys()) {
		m.Delete(key)
		_, ok := m.Load(key)
		if ok {
			t.Fatal("found after delete")
		}
	}
}

type fakeint int

func (a fakeint) Compare(b fakeint) int {
	return cmp.Compare(a, b)
}

func genfakeint() fakeint {
	return rand.N[fakeint](math.MaxInt)
}

func TestBSMapT(t *testing.T) {
	m := NewBSMapT[fakeint, Empty]()
	for range 100 {
		m.Store(genfakeint(), Empty{})
	}
	if !slices.IsSorted(m.SliceKeys()) {
		t.Fatal("not sorted")
	}
	for key := range m.Keys {
		_, ok := m.Load(key)
		if !ok {
			t.Fatal("not found")
		}
	}
	for _, key := range cslices.Random(m.SliceKeys()) {
		m.Delete(key)
		_, ok := m.Load(key)
		if ok {
			t.Fatal("found after delete")
		}
	}
}

func BenchmarkBSMap(b *testing.B) {
	m := NewBSMap(make(PairSlice[fakeint, Empty], 0, b.N)...)
	for range b.N {
		m.Store(genfakeint(), Empty{})
	}
}

func BenchmarkBSMapT(b *testing.B) {
	m := NewBSMapT(make(PairSlice[fakeint, Empty], 0, b.N)...)
	for range b.N {
		m.Store(genfakeint(), Empty{})
	}
}

func BenchmarkBuiltinMap(b *testing.B) {
	m := make(map[fakeint]Empty, b.N)
	for range b.N {
		m[genfakeint()] = Empty{}
	}
}

func TestClear(t *testing.T) {
	m := NewBSMap[int, string]()
	m.Store(1, "1")
	m.Clear()
	println(m)
}
