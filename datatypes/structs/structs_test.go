package cstructs

import (
	"fmt"
	"testing"
)

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
	m, err := ToMap(a)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(m)
}
