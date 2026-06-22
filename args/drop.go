package args

func Drop1[T, U any](t T, _ U) T {
	return t
}
