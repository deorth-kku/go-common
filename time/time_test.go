package ctime

import (
	"testing"
	"time"
)

func BenchmarkIsZero(b *testing.B) {
	var a time.Time
	for range b.N {
		_ = a.IsZero()
	}
}
