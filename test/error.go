package ctest

import (
	"errors"
)

func NoError[T TestingCommon](t T, err error, customMsg ...any) {
	if err == nil {
		return
	}
	t.Helper()
	if len(customMsg) == 0 {
		t.Fatal(err)
	} else {
		t.Fatal(customMsg...)

	}
}

func IsError[T TestingCommon](t T, actual, expected error, customMsg ...any) {
	if errors.Is(actual, expected) {
		return
	}
	t.Helper()
	if len(customMsg) == 0 {
		t.Fatal(actual)
	} else {
		t.Fatal(customMsg...)
	}
}

func Error[T TestingCommon](t T, err error, customMsg ...any) {
	if err != nil {
		return
	}
	t.Helper()
	if len(customMsg) == 0 {
		t.Fatal("no error when expecting an error")
	} else {
		t.Fatal(customMsg...)
	}
}

func AsErrorType[E error, T TestingCommon](t T, err error, customMsg ...any) E {
	var e E
	if errors.As(err, &e) {
		return e
	}
	t.Helper()
	if len(customMsg) == 0 {
		t.Fatal(err)
	} else {
		t.Fatal(customMsg...)
	}
	return e
}
