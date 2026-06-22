package ccmp

import (
	"reflect"
	"testing"
	"time"
)

func BenchmarkReflectIsZero(b *testing.B) {
	var t time.Time
	for range b.N {
		reflect.ValueOf(t).IsZero()
	}
}

func BenchmarkGenericsIsZero(b *testing.B) {
	var t time.Time
	for range b.N {
		IsZero(t)
	}
}
