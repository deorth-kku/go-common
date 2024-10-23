package common

import (
	"fmt"
	"log/slog"
	"testing"
)

func TestHttp(t *testing.T) {
	server := NewHttpServer()
	server.ListenAndServe("/tmp/123.sock,0666")
}

func TestCutSlice(t *testing.T) {
	longslice := make([]int, 65535)
	for i := range longslice {
		longslice[i] = i
	}
	last := -1
	for _, subslice := range CutSlice(longslice, 100) {
		if subslice[0] != last+1 {
			t.Error("no!")
		}
		last = subslice[len(subslice)-1]
	}
}

func TestNaN32(t *testing.T) {
	f := Nan32()
	if !IsNaN(f) {
		t.Error("no!")
	}
}

func TestSetLog(t *testing.T) {
	closer, err := SetLog("", "DEBUG", SlogHideTime{}, SlogJson{})
	if err != nil {
		t.Error(err)
	}
	defer closer.Close()
	slog.Debug("test")
}

func TestSetGroup(t *testing.T) {
	closer, err := SetLog("", "DEBUG", SlogHideTime{})
	if err != nil {
		t.Error(err)
	}
	defer closer.Close()
	slog.Info("test", "test", Map2Group(map[string]any{
		"a": 1,
		"b": 2,
		"g": map[string]any{
			"a": 1,
			"b": 2,
		},
	}))

}

func TestStruct(t *testing.T) {
	a := struct {
		A int
		B string
		M map[string]any
	}{
		A: 1,
		B: "2",
		M: map[string]any{
			"test": 1,
		},
	}
	m, err := Struct2Map(a)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(m)
}
