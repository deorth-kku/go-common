package common

import (
	"fmt"
	"log/slog"
	"maps"
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
	err := SetLog("", "DEBUG", "JSON", SlogHideTime{})
	if err != nil {
		t.Error(err)
	}

	slog.Debug("test")
}

func TestSetGroup(t *testing.T) {
	err := SetLog("", "DEBUG", "DEFAULT", SlogHideTime{})
	if err != nil {
		t.Error(err)
	}
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
	a := mix{
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

func TestSlogIterMap(t *testing.T) {
	m := map[string]any{
		"a": 1,
		"b": 2,
	}
	err := SetLog("", "DEBUG", "DEFAULT", SlogHideTime{}, SlogIter{}, SlogMap{})
	if err != nil {
		t.Error(err)
	}

	slog.Info("test log iter and map", "iter", maps.All(m), "map", m)
}

type mix struct {
	A int
	B string
	M map[string]any
}

func TestSlogStruct(t *testing.T) {
	a := mix{
		A: 1,
		B: "2",
		M: map[string]any{
			"test": 1,
		},
	}
	err := SetLog("", "DEBUG", "DEFAULT", SlogStruct[mix]{}, SlogHideTime{}, SlogIter{}, SlogMap{})
	if err != nil {
		t.Error(err)
	}
	slog.Info("test struct", "a", a)
}

func TestParseMode(t *testing.T) {
	_, m, err := FileWithMode("test,0666")
	if err != nil {
		t.Error(err)
		return
	}
	if m != 0666 {
		t.Error("wrong")
	}
}
