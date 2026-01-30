package ctest

import (
	"errors"
	"testing"
)

func NoError(t *testing.T, err error, customMsg ...any) {
	if err == nil {
		return
	}
	if len(customMsg) == 0 {
		t.Fatal(customMsg...)
	} else {
		t.Fatal(err)
	}
}

func IsError(t *testing.T, actual, expected error, customMsg ...any) {
	if errors.Is(actual, expected) {
		return
	}
	if len(customMsg) == 0 {
		t.Fatal(customMsg...)
	} else {
		t.Fatal(actual)
	}
}

func Error(t *testing.T, err error, customMsg ...any) {
	if err != nil {
		return
	}
	if len(customMsg) == 0 {
		t.Fatal("no error when expecting an error")
	} else {
		t.Fatal(customMsg...)
	}
}

func AsErrorType[T error](t *testing.T, err error, customMsg ...any) T {
	var e T
	if errors.As(err, &e) {
		return e
	}
	if len(customMsg) == 0 {
		t.Fatal(err)
	} else {
		t.Fatal(customMsg...)
	}
	return e
}
