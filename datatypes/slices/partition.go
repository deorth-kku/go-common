package cslices

func PartitionInPlace[E any, S ~[]E](s S, f func(E) bool) (true, false S) {
	if s == nil {
		return nil, nil
	}

	i, t := 0, len(s)-1
	for i <= t {
		if f(s[i]) {
			i++
			continue
		}

		s[i], s[t] = s[t], s[i]
		t--
	}
	return s[:i], s[i:]
}
