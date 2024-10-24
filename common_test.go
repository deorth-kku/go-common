package common

import (
	"fmt"
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

type mix struct {
	A int
	B string
	M map[string]any
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
