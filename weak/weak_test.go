package cweak

import (
	"fmt"
	"net/url"
	"testing"
)

func BenchmarkWeak(b *testing.B) {
	cli := new(url.URL)
	for b.Loop() {
		w := NewInterface[fmt.Stringer](cli)
		_ = w
	}
}

func BenchmarkWeakValue(b *testing.B) {
	cli := new(url.URL)
	w := NewInterface[fmt.Stringer](cli)
	for b.Loop() {
		w.Value()
	}
}

type isAbs interface {
	IsAbs() bool
}

func BenchmarkTypeAssert(b *testing.B) {
	cli := new(url.URL)
	w := NewInterface[fmt.Stringer](cli)
	for b.Loop() {
		_, _ = TypeAssert[isAbs](w)
	}
}

func BenchmarkTypeAssertSlow(b *testing.B) {
	cli := new(url.URL)
	w := NewInterface[fmt.Stringer](cli)
	for b.Loop() {
		_, _ = w.Value().(isAbs)
	}
}
