package ctest

import (
	"cmp"
	"reflect"
)

func Equal[C comparable, T TestingCommon](t T, a1, a2 C, customMsg ...any) {
	if a1 == a2 {
		return
	}
	t.Helper()
	if len(customMsg) == 0 {
		t.Fatalf("%v is not equal to %v", a1, a2)
	} else {
		t.Fatal(customMsg...)
	}
}

func DeepEqual[T TestingCommon](t T, a1, a2 any, customMsg ...any) {
	if reflect.DeepEqual(a1, a2) {
		return
	}
	t.Helper()
	if len(customMsg) == 0 {
		t.Fatalf("%v is not equal to %v", a1, a2)
	} else {
		t.Fatal(customMsg...)
	}
}

func Less[O cmp.Ordered, T TestingCommon](t T, a1, a2 O, customMsg ...any) {
	if a1 < a2 {
		return
	}
	t.Helper()
	if len(customMsg) == 0 {
		t.Fatalf("%v is not less than %v", a1, a2)
	} else {
		t.Fatal(customMsg...)
	}
}

func Greater[O cmp.Ordered, T TestingCommon](t T, a1, a2 O, customMsg ...any) {
	if a1 > a2 {
		return
	}
	t.Helper()
	if len(customMsg) == 0 {
		t.Fatalf("%v is not greater than %v", a1, a2)
	} else {
		t.Fatal(customMsg...)
	}
}

func LessOrEqual[O cmp.Ordered, T TestingCommon](t T, a1, a2 O, customMsg ...any) {
	if a1 <= a2 {
		return
	}
	t.Helper()
	if len(customMsg) == 0 {
		t.Fatalf("%v is not less than or equal to %v", a1, a2)
	} else {
		t.Fatal(customMsg...)
	}
}

func GreaterOrEqual[O cmp.Ordered, T TestingCommon](t T, a1, a2 O, customMsg ...any) {
	if a1 >= a2 {
		return
	}
	t.Helper()
	if len(customMsg) == 0 {
		t.Fatalf("%v is not greater than or equal to %v", a1, a2)
	} else {
		t.Fatal(customMsg...)
	}
}
