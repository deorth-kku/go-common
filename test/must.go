package ctest

func Must00[T TestingCommon, F ~func() error](t T, f F) {
	err := f()
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
}

func Must01[T TestingCommon, F ~func() (R1, error), R1 any](t T, f F) R1 {
	r1, err := f()
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
	return r1
}

func Must10[T TestingCommon, F ~func(A1) error, A1 any](t T, f F, a1 A1) {
	err := f(a1)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
}

func Must11[T TestingCommon, F ~func(A1) (R1, error), A1, R1 any](t T, f F, a1 A1) R1 {
	r1, err := f(a1)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
	return r1
}

func Must12[T TestingCommon, F ~func(A1, A2) error, A1, A2 any](t T, f F, a1 A1, a2 A2) {
	err := f(a1, a2)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
}

func Must20[T TestingCommon, F ~func(A1, A2) error, A1, A2 any](t T, f F, a1 A1, a2 A2) {
	err := f(a1, a2)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
}

func Must21[T TestingCommon, F ~func(A1, A2) (R1, error), A1, A2, R1 any](t T, f F, a1 A1, a2 A2) R1 {
	r1, err := f(a1, a2)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
	return r1
}

func Must22[T TestingCommon, F ~func(A1, A2) (R1, R2, error), A1, A2, R1, R2 any](t T, f F, a1 A1, a2 A2) (R1, R2) {
	r1, r2, err := f(a1, a2)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
	return r1, r2
}

func Must23[T TestingCommon, F ~func(A1, A2, A3) error, A1, A2, A3 any](t T, f F, a1 A1, a2 A2, a3 A3) {
	err := f(a1, a2, a3)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
}

func Must30[T TestingCommon, F ~func(A1, A2, A3) error, A1, A2, A3 any](t T, f F, a1 A1, a2 A2, a3 A3) {
	err := f(a1, a2, a3)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
}

func Must31[T TestingCommon, F ~func(A1, A2, A3) (R1, error), A1, A2, A3, R1 any](t T, f F, a1 A1, a2 A2, a3 A3) R1 {
	r1, err := f(a1, a2, a3)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
	return r1
}

func Must32[T TestingCommon, F ~func(A1, A2, A3) (R1, R2, error), A1, A2, A3, R1, R2 any](t T, f F, a1 A1, a2 A2, a3 A3) (R1, R2) {
	r1, r2, err := f(a1, a2, a3)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
	return r1, r2
}

func Must33[T TestingCommon, F ~func(A1, A2, A3) (R1, R2, R3, error), A1, A2, A3, R1, R2, R3 any](t T, f F, a1 A1, a2 A2, a3 A3) (R1, R2, R3) {
	r1, r2, r3, err := f(a1, a2, a3)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
	return r1, r2, r3
}

func Must40[T TestingCommon, F ~func(A1, A2, A3, A4) error, A1, A2, A3, A4 any](t T, f F, a1 A1, a2 A2, a3 A3, a4 A4) {
	err := f(a1, a2, a3, a4)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
}

func Must41[T TestingCommon, F ~func(A1, A2, A3, A4) (R1, error), A1, A2, A3, A4, R1 any](t T, f F, a1 A1, a2 A2, a3 A3, a4 A4) R1 {
	r1, err := f(a1, a2, a3, a4)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
	return r1
}

func Must42[T TestingCommon, F ~func(A1, A2, A3, A4) (R1, R2, error), A1, A2, A3, A4, R1, R2 any](t T, f F, a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2) {
	r1, r2, err := f(a1, a2, a3, a4)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
	return r1, r2
}

func Must43[T TestingCommon, F ~func(A1, A2, A3, A4) (R1, R2, R3, error), A1, A2, A3, A4, R1, R2, R3 any](t T, f F, a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2, R3) {
	r1, r2, r3, err := f(a1, a2, a3, a4)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
	return r1, r2, r3
}

func Must44[T TestingCommon, F ~func(A1, A2, A3, A4) (R1, R2, R3, R4, error), A1, A2, A3, A4, R1, R2, R3, R4 any](t T, f F, a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2, R3, R4) {
	r1, r2, r3, r4, err := f(a1, a2, a3, a4)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
	return r1, r2, r3, r4
}

func Must50[T TestingCommon, F ~func(A1, A2, A3, A4, A5) error, A1, A2, A3, A4, A5 any](t T, f F, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) {
	err := f(a1, a2, a3, a4, a5)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
}

func Must51[T TestingCommon, F ~func(A1, A2, A3, A4, A5) (R1, error), A1, A2, A3, A4, A5, R1 any](t T, f F, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) R1 {
	r1, err := f(a1, a2, a3, a4, a5)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
	return r1
}

func Must52[T TestingCommon, F ~func(A1, A2, A3, A4, A5) (R1, R2, error), A1, A2, A3, A4, A5, R1, R2 any](t T, f F, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2) {
	r1, r2, err := f(a1, a2, a3, a4, a5)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
	return r1, r2
}

func Must53[T TestingCommon, F ~func(A1, A2, A3, A4, A5) (R1, R2, R3, error), A1, A2, A3, A4, A5, R1, R2, R3 any](t T, f F, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2, R3) {
	r1, r2, r3, err := f(a1, a2, a3, a4, a5)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
	return r1, r2, r3
}

func Must00Ex[T TestingCommon, F ~func(...AEx) error, AEx any](t T, f F, aex ...AEx) {
	err := f(aex...)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
}

func Must01Ex[T TestingCommon, F ~func(...AEx) (R1, error), AEx, R1 any](t T, f F, aex ...AEx) R1 {
	r1, err := f(aex...)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
	return r1
}

func Must10Ex[T TestingCommon, F ~func(A1, ...AEx) error, A1, AEx any](t T, f F, a1 A1, aex ...AEx) {
	err := f(a1, aex...)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
}

func Must11Ex[T TestingCommon, F ~func(A1, ...AEx) (R1, error), A1, AEx, R1 any](t T, f F, a1 A1, aex ...AEx) R1 {
	r1, err := f(a1, aex...)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
	return r1
}

func Must12Ex[T TestingCommon, F ~func(A1, A2, ...AEx) error, A1, A2, AEx any](t T, f F, a1 A1, a2 A2, aex ...AEx) {
	err := f(a1, a2, aex...)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
}

func Must20Ex[T TestingCommon, F ~func(A1, A2, ...AEx) error, A1, A2, AEx any](t T, f F, a1 A1, a2 A2, aex ...AEx) {
	err := f(a1, a2, aex...)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
}

func Must21Ex[T TestingCommon, F ~func(A1, A2, ...AEx) (R1, error), A1, A2, AEx, R1 any](t T, f F, a1 A1, a2 A2, aex ...AEx) R1 {
	r1, err := f(a1, a2, aex...)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
	return r1
}

func Must22Ex[T TestingCommon, F ~func(A1, A2, ...AEx) (R1, R2, error), A1, A2, AEx, R1, R2 any](t T, f F, a1 A1, a2 A2, aex ...AEx) (R1, R2) {
	r1, r2, err := f(a1, a2, aex...)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
	return r1, r2
}

func Must23Ex[T TestingCommon, F ~func(A1, A2, A3, ...AEx) error, A1, A2, A3, AEx any](t T, f F, a1 A1, a2 A2, a3 A3, aex ...AEx) {
	err := f(a1, a2, a3, aex...)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
}

func Must30Ex[T TestingCommon, F ~func(A1, A2, A3, ...AEx) error, A1, A2, A3, AEx any](t T, f F, a1 A1, a2 A2, a3 A3, aex ...AEx) {
	err := f(a1, a2, a3, aex...)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
}

func Must31Ex[T TestingCommon, F ~func(A1, A2, A3, ...AEx) (R1, error), A1, A2, A3, AEx, R1 any](t T, f F, a1 A1, a2 A2, a3 A3, aex ...AEx) R1 {
	r1, err := f(a1, a2, a3, aex...)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
	return r1
}

func Must32Ex[T TestingCommon, F ~func(A1, A2, A3, ...AEx) (R1, R2, error), A1, A2, A3, AEx, R1, R2 any](t T, f F, a1 A1, a2 A2, a3 A3, aex ...AEx) (R1, R2) {
	r1, r2, err := f(a1, a2, a3, aex...)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
	return r1, r2
}

func Must33Ex[T TestingCommon, F ~func(A1, A2, A3, ...AEx) (R1, R2, R3, error), A1, A2, A3, AEx, R1, R2, R3 any](t T, f F, a1 A1, a2 A2, a3 A3, aex ...AEx) (R1, R2, R3) {
	r1, r2, r3, err := f(a1, a2, a3, aex...)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
	return r1, r2, r3
}

func Must40Ex[T TestingCommon, F ~func(A1, A2, A3, A4, ...AEx) error, A1, A2, A3, A4, AEx any](t T, f F, a1 A1, a2 A2, a3 A3, a4 A4, aex ...AEx) {
	err := f(a1, a2, a3, a4, aex...)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
}

func Must41Ex[T TestingCommon, F ~func(A1, A2, A3, A4, ...AEx) (R1, error), A1, A2, A3, A4, AEx, R1 any](t T, f F, a1 A1, a2 A2, a3 A3, a4 A4, aex ...AEx) R1 {
	r1, err := f(a1, a2, a3, a4, aex...)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
	return r1
}

func Must42Ex[T TestingCommon, F ~func(A1, A2, A3, A4, ...AEx) (R1, R2, error), A1, A2, A3, A4, AEx, R1, R2 any](t T, f F, a1 A1, a2 A2, a3 A3, a4 A4, aex ...AEx) (R1, R2) {
	r1, r2, err := f(a1, a2, a3, a4, aex...)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
	return r1, r2
}

func Must43Ex[T TestingCommon, F ~func(A1, A2, A3, A4, ...AEx) (R1, R2, R3, error), A1, A2, A3, A4, AEx, R1, R2, R3 any](t T, f F, a1 A1, a2 A2, a3 A3, a4 A4, aex ...AEx) (R1, R2, R3) {
	r1, r2, r3, err := f(a1, a2, a3, a4, aex...)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
	return r1, r2, r3
}

func Must44Ex[T TestingCommon, F ~func(A1, A2, A3, A4, ...AEx) (R1, R2, R3, R4, error), A1, A2, A3, A4, AEx, R1, R2, R3, R4 any](t T, f F, a1 A1, a2 A2, a3 A3, a4 A4, aex ...AEx) (R1, R2, R3, R4) {
	r1, r2, r3, r4, err := f(a1, a2, a3, a4, aex...)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
	return r1, r2, r3, r4
}

func Must50Ex[T TestingCommon, F ~func(A1, A2, A3, A4, A5, ...AEx) error, A1, A2, A3, A4, A5, AEx any](t T, f F, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, aex ...AEx) {
	err := f(a1, a2, a3, a4, a5, aex...)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
}

func Must51Ex[T TestingCommon, F ~func(A1, A2, A3, A4, A5, ...AEx) (R1, error), A1, A2, A3, A4, A5, AEx, R1 any](t T, f F, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, aex ...AEx) R1 {
	r1, err := f(a1, a2, a3, a4, a5, aex...)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
	return r1
}

func Must52Ex[T TestingCommon, F ~func(A1, A2, A3, A4, A5, ...AEx) (R1, R2, error), A1, A2, A3, A4, A5, AEx, R1, R2 any](t T, f F, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, aex ...AEx) (R1, R2) {
	r1, r2, err := f(a1, a2, a3, a4, a5, aex...)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
	return r1, r2
}

func Must53Ex[T TestingCommon, F ~func(A1, A2, A3, A4, A5, ...AEx) (R1, R2, R3, error), A1, A2, A3, A4, A5, AEx, R1, R2, R3 any](t T, f F, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, aex ...AEx) (R1, R2, R3) {
	r1, r2, r3, err := f(a1, a2, a3, a4, a5, aex...)
	if err != nil {
		t.Helper()
		t.Fatal(err)
	}
	return r1, r2, r3
}
