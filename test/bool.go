package ctest

func True[T TestingCommon](t T, b bool, customMsg ...any) {
	if b {
		return
	}
	t.Helper()
	if len(customMsg) == 0 {
		t.Fatal("getting false when require true")
	} else {
		t.Fatal(customMsg...)
	}
}

func False[T TestingCommon](t T, b bool, customMsg ...any) {
	if !b {
		return
	}
	t.Helper()
	if len(customMsg) == 0 {
		t.Fatal("getting true when require false")
	} else {
		t.Fatal(customMsg...)
	}
}

func True00[T TestingCommon, F ~func() bool](t T, f F, customMsg ...any) {
	t.Helper()
	True(t, f(), customMsg...)
}

func True01[T TestingCommon, F ~func() (R1, bool), R1 any](t T, f F, customMsg ...any) R1 {
	v, ok := f()
	t.Helper()
	True(t, ok, customMsg...)
	return v
}

func True02[T TestingCommon, F ~func() (R1, R2, bool), R1, R2 any](t T, f F, customMsg ...any) (R1, R2) {
	v1, v2, ok := f()
	t.Helper()
	True(t, ok, customMsg...)
	return v1, v2
}

func True03[T TestingCommon, F ~func() (R1, R2, R3, bool), R1, R2, R3 any](t T, f F, customMsg ...any) (R1, R2, R3) {
	v1, v2, v3, ok := f()
	t.Helper()
	True(t, ok, customMsg...)
	return v1, v2, v3
}

func True04[T TestingCommon, F ~func() (R1, R2, R3, R4, bool), R1, R2, R3, R4 any](t T, f F, customMsg ...any) (R1, R2, R3, R4) {
	v1, v2, v3, v4, ok := f()
	t.Helper()
	True(t, ok, customMsg...)
	return v1, v2, v3, v4
}

func True05[T TestingCommon, F ~func() (R1, R2, R3, R4, R5, bool), R1, R2, R3, R4, R5 any](t T, f F, customMsg ...any) (R1, R2, R3, R4, R5) {
	v1, v2, v3, v4, v5, ok := f()
	t.Helper()
	True(t, ok, customMsg...)
	return v1, v2, v3, v4, v5
}

func True06[T TestingCommon, F ~func() (R1, R2, R3, R4, R5, R6, bool), R1, R2, R3, R4, R5, R6 any](t T, f F, customMsg ...any) (R1, R2, R3, R4, R5, R6) {
	v1, v2, v3, v4, v5, v6, ok := f()
	t.Helper()
	True(t, ok, customMsg...)
	return v1, v2, v3, v4, v5, v6
}

func True07[T TestingCommon, F ~func() (R1, R2, R3, R4, R5, R6, R7, bool), R1, R2, R3, R4, R5, R6, R7 any](t T, f F, customMsg ...any) (R1, R2, R3, R4, R5, R6, R7) {
	v1, v2, v3, v4, v5, v6, v7, ok := f()
	t.Helper()
	True(t, ok, customMsg...)
	return v1, v2, v3, v4, v5, v6, v7
}

func True08[T TestingCommon, F ~func() (R1, R2, R3, R4, R5, R6, R7, R8, bool), R1, R2, R3, R4, R5, R6, R7, R8 any](t T, f F, customMsg ...any) (R1, R2, R3, R4, R5, R6, R7, R8) {
	v1, v2, v3, v4, v5, v6, v7, v8, ok := f()
	t.Helper()
	True(t, ok, customMsg...)
	return v1, v2, v3, v4, v5, v6, v7, v8
}

func True09[T TestingCommon, F ~func() (R1, R2, R3, R4, R5, R6, R7, R8, R9, bool), R1, R2, R3, R4, R5, R6, R7, R8, R9 any](t T, f F, customMsg ...any) (R1, R2, R3, R4, R5, R6, R7, R8, R9) {
	v1, v2, v3, v4, v5, v6, v7, v8, v9, ok := f()
	t.Helper()
	True(t, ok, customMsg...)
	return v1, v2, v3, v4, v5, v6, v7, v8, v9
}

func True10[T TestingCommon, F ~func(A1) bool, A1 any](t T, f F, a1 A1, customMsg ...any) {
	t.Helper()
	True(t, f(a1), customMsg...)
}

func True11[T TestingCommon, F ~func(A1) (R1, bool), A1, R1 any](t T, f F, a1 A1, customMsg ...any) R1 {
	v, ok := f(a1)
	t.Helper()
	True(t, ok, customMsg...)
	return v
}

func True12[T TestingCommon, F ~func(A1) (R1, R2, bool), A1, R1, R2 any](t T, f F, a1 A1, customMsg ...any) (R1, R2) {
	v1, v2, ok := f(a1)
	t.Helper()
	True(t, ok, customMsg...)
	return v1, v2
}

func True13[T TestingCommon, F ~func(A1) (R1, R2, R3, bool), A1, R1, R2, R3 any](t T, f F, a1 A1, customMsg ...any) (R1, R2, R3) {
	v1, v2, v3, ok := f(a1)
	t.Helper()
	True(t, ok, customMsg...)
	return v1, v2, v3
}

func True14[T TestingCommon, F ~func(A1) (R1, R2, R3, R4, bool), A1, R1, R2, R3, R4 any](t T, f F, a1 A1, customMsg ...any) (R1, R2, R3, R4) {
	v1, v2, v3, v4, ok := f(a1)
	t.Helper()
	True(t, ok, customMsg...)
	return v1, v2, v3, v4
}

func True15[T TestingCommon, F ~func(A1) (R1, R2, R3, R4, R5, bool), A1, R1, R2, R3, R4, R5 any](t T, f F, a1 A1, customMsg ...any) (R1, R2, R3, R4, R5) {
	v1, v2, v3, v4, v5, ok := f(a1)
	t.Helper()
	True(t, ok, customMsg...)
	return v1, v2, v3, v4, v5
}

func True20[T TestingCommon, F ~func(A1, A2) bool, A1, A2 any](t T, f F, a1 A1, a2 A2, customMsg ...any) {
	t.Helper()
	True(t, f(a1, a2), customMsg...)
}

func True21[T TestingCommon, F ~func(A1, A2) (R1, bool), A1, A2, R1 any](t T, f F, a1 A1, a2 A2, customMsg ...any) R1 {
	v, ok := f(a1, a2)
	t.Helper()
	True(t, ok, customMsg...)
	return v
}

func True22[T TestingCommon, F ~func(A1, A2) (R1, R2, bool), A1, A2, R1, R2 any](t T, f F, a1 A1, a2 A2, customMsg ...any) (R1, R2) {
	v1, v2, ok := f(a1, a2)
	t.Helper()
	True(t, ok, customMsg...)
	return v1, v2
}

func True23[T TestingCommon, F ~func(A1, A2) (R1, R2, R3, bool), A1, A2, R1, R2, R3 any](t T, f F, a1 A1, a2 A2, customMsg ...any) (R1, R2, R3) {
	v1, v2, v3, ok := f(a1, a2)
	t.Helper()
	True(t, ok, customMsg...)
	return v1, v2, v3
}

func True24[T TestingCommon, F ~func(A1, A2) (R1, R2, R3, R4, bool), A1, A2, R1, R2, R3, R4 any](t T, f F, a1 A1, a2 A2, customMsg ...any) (R1, R2, R3, R4) {
	v1, v2, v3, v4, ok := f(a1, a2)
	t.Helper()
	True(t, ok, customMsg...)
	return v1, v2, v3, v4
}

func True30[T TestingCommon, F ~func(A1, A2, A3) bool, A1, A2, A3 any](t T, f F, a1 A1, a2 A2, a3 A3, customMsg ...any) {
	t.Helper()
	True(t, f(a1, a2, a3), customMsg...)
}

func True31[T TestingCommon, F ~func(A1, A2, A3) (R1, bool), A1, A2, A3, R1 any](t T, f F, a1 A1, a2 A2, a3 A3, customMsg ...any) R1 {
	v, ok := f(a1, a2, a3)
	t.Helper()
	True(t, ok, customMsg...)
	return v
}

func True32[T TestingCommon, F ~func(A1, A2, A3) (R1, R2, bool), A1, A2, A3, R1, R2 any](t T, f F, a1 A1, a2 A2, a3 A3, customMsg ...any) (R1, R2) {
	v1, v2, ok := f(a1, a2, a3)
	t.Helper()
	True(t, ok, customMsg...)
	return v1, v2
}
