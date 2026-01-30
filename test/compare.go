package ctest

import (
	"cmp"
	"reflect"
	"testing"
)

func Equal[T comparable](t *testing.T, a1, a2 T, customMsg ...any) {
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

func DeepEqual(t *testing.T, a1, a2 any, customMsg ...any) {
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

func Less[T cmp.Ordered](t *testing.T, a1, a2 T, customMsg ...any) {
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

func Greater[T cmp.Ordered](t *testing.T, a1, a2 T, customMsg ...any) {
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

func LessOrEqual[T cmp.Ordered](t *testing.T, a1, a2 T, customMsg ...any) {
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

func GreaterOrEqual[T cmp.Ordered](t *testing.T, a1, a2 T, customMsg ...any) {
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
