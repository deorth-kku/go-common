package common

func AnySliceReplaceNil(in []string) (out []any) {
	out = make([]any, len(in))
	for i, arg := range in {
		if len(arg) == 0 {
			out[i] = nil
		} else {
			out[i] = arg
		}
	}
	return
}

func AnySlice[t any](in []t) (out []any) {
	out = make([]any, len(in))
	for i, arg := range in {
		out[i] = arg
	}
	return
}

func CutSlice[t any](in []t, l int) (out [][]t) {
	rounds := len(in) / l
	out = make([][]t, rounds+1)

	for i := 0; i < rounds; i++ {
		out[i] = in[i*l : (i+1)*l]
	}
	out[rounds] = in[rounds*l:]
	return
}
