package common

type Set[T comparable] struct {
	data map[T]struct{}
}

func NewSet[T comparable](size ...int) (s Set[T]) {
	if len(size) > 0 {
		s.data = make(map[T]struct{}, size[0])
	} else {
		s.data = make(map[T]struct{})
	}
	return
}

func NewSetFromSlice[T comparable](slice []T) (s Set[T]) {
	s.data = make(map[T]struct{}, len(slice))
	for _, elem := range slice {
		s.Add(elem)
	}
	return
}

func (s Set[T]) Add(elem T) {
	s.data[elem] = struct{}{}
}

func (s Set[T]) Delete(elem T) {
	delete(s.data, elem)
}

func (s Set[T]) Len() int {
	return len(s.data)
}

func (s Set[T]) Has(elem T) (ok bool) {
	_, ok = s.data[elem]
	return
}

func (s Set[T]) Range(yield func(T) bool) {
	for key := range s.data {
		if !yield(key) {
			return
		}
	}
}

func (s Set[T]) Slice() (slice []T) {
	slice = make([]T, 0, s.Len())
	for elem := range s.data {
		slice = append(slice, elem)
	}
	return
}
