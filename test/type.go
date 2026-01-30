package ctest

import (
	"reflect"
	"testing"
)

func IsType[T any](t *testing.T, arg any, customMsg ...any) T {
	v, ok := arg.(T)
	if ok {
		return v
	}
	if len(customMsg) == 0 {
		t.Fatalf("%v is not a %s type", arg, reflect.TypeFor[T]())
	} else {
		t.Fatal(customMsg...)
	}
	return v
}

func IsReflectType(t *testing.T, arg any, ty reflect.Type, customMsg ...any) {
	if reflect.TypeOf(arg) == ty {
		return
	}
	if len(customMsg) == 0 {
		t.Fatalf("%v is not a %s type", arg, ty)
	} else {
		t.Fatal(customMsg...)
	}
}
