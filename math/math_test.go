package cmath

import (
	"math"
	"math/rand/v2"
	"testing"
)

func TestNaN32(t *testing.T) {
	f := NaN32()
	if !IsNaN(f) {
		t.Error("no!")
	}
}

func TestInf32(t *testing.T) {
	f := Inf32(1)
	if !IsInf(f, 1) {
		t.Error("no!")
	}
}

func TestParse(t *testing.T) {
	type ii int
	str := "-1"
	i, err := Parse[ii](str, 10)
	if err != nil {
		t.Error(err)
		return
	}
	if str != Format(ii(i), 10) {
		t.Error("not match")
		return
	}

	type uu uint16
	str = "443"
	u8, err := Parse[uu](str, 10)
	if err != nil {
		t.Error(err)
		return
	}
	if str != Format(uu(u8), 10) {
		t.Error("not match")
		return
	}
	_, err = Parse[uu]("65536", 10)
	if err == nil {
		t.Error("not overflow when it should")
		return
	}

	str = "1.2345"
	f, err := Parse[float32](str, 0)
	if err != nil {
		t.Error(err)
		return
	}
	if str != Format(f, 0) {
		t.Error("not match")
		return
	}

	f64, err := Parse[float64](str, 0)
	if err != nil {
		t.Error(err)
		return
	}
	if str != Format(f64, 0) {
		t.Error("not match")
		return
	}
}

func TestMaxInt(t *testing.T) {
	type uu uint
	if MaxInt[uu]() != math.MaxUint {
		t.Error("not eq")
	}
	type ii int
	if MaxInt[ii]() != math.MaxInt {
		t.Error("not eq")
	}
}

func BenchmarkInf(b *testing.B) {
	num := math.Inf(1)
	for range b.N {
		IsInf(num, 1)
	}
}

func BenchmarkInfStd(b *testing.B) {
	num := math.Inf(1)
	for range b.N {
		math.IsInf(num, 1)
	}
}

func BenchmarkInf32(b *testing.B) {
	num := Inf32(1)
	for range b.N {
		IsInf(num, 1)
	}
}

func BenchmarkInf32Std(b *testing.B) {
	num := Inf32(1)
	for range b.N {
		math.IsInf(float64(num), 1)
	}
}

func BenchmarkCeil(b *testing.B) {
	for range b.N {
		DevidedCeil(rand.Int(), rand.Int())
	}
}
