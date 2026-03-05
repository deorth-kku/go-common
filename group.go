package common

func GroupParse[I, O any](parse func(i I) (O, error), ins []I) (os []O, err error) {
	if ins == nil {
		return nil, nil
	}
	os = make([]O, len(ins))
	for i, in := range ins {
		os[i], err = parse(in)
		if err != nil {
			return nil, err
		}
	}
	return
}

func GroupConv[I, O any](parse func(i I) O, ins []I) (os []O) {
	if ins == nil {
		return
	}
	os = make([]O, len(ins))
	for i, in := range ins {
		os[i] = parse(in)
	}
	return
}

func GroupGet[I, O any](get func(i I) (O, bool), ins []I) (os []O, ok bool) {
	if ins == nil {
		return nil, true
	}
	os = make([]O, len(ins))
	for i, in := range ins {
		os[i], ok = get(in)
		if !ok {
			return nil, ok
		}
	}
	return os, true
}
