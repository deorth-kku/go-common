package ctest

import (
	"reflect"
)

func IsType[T TestingCommon, V any](t T, arg any, customMsg ...any) V {
	v, ok := arg.(V)
	if ok {
		return v
	}
	t.Helper()
	if len(customMsg) == 0 {
		t.Fatalf("%v is not a %s type", arg, reflect.TypeFor[V]())
	} else {
		t.Fatal(customMsg...)
	}
	return v
}

func IsReflectType[T TestingCommon](t T, arg any, ty reflect.Type, customMsg ...any) {
	if reflect.TypeOf(arg) == ty {
		return
	}
	t.Helper()
	if len(customMsg) == 0 {
		t.Fatalf("%v is not a %s type", arg, ty)
	} else {
		t.Fatal(customMsg...)
	}
}
