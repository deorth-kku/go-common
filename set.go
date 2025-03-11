package common

import "maps"

type Empty = struct{}

type Set[T comparable] struct {
	data map[T]Empty
}

func NewSet[T comparable](size ...int) (s Set[T]) {
	if len(size) > 0 {
		s.data = make(map[T]Empty, size[0])
	} else {
		s.data = make(map[T]Empty)
	}
	return
}

func NewSetFromSlice[T comparable](slice []T) (s Set[T]) {
	s.data = make(map[T]Empty, len(slice))
	for _, elem := range slice {
		s.Add(elem)
	}
	return
}

func (s *Set[T]) Add(elem T) {
	if s.data == nil {
		s.data = make(map[T]Empty)
	}
	s.data[elem] = Empty{}
}

func (s Set[T]) Delete(elem T) {
	delete(s.data, elem)
}

func (s Set[T]) Len() int {
	return len(s.data)
}

func (s Set[T]) Has(elem T) (ok bool) {
	if s.data == nil {
		return false
	}
	_, ok = s.data[elem]
	return
}

func (s Set[T]) Range(yield Yield[T]) {
	if s.data == nil {
		return
	}
	maps.Keys(s.data)(yield)
}

func (s Set[T]) Slice() []T {
	if s.data == nil {
		return nil
	}
	return SliceCollect(s.Range, len(s.data))
}

func (s Set[T]) Clone() Set[T] {
	return Set[T]{maps.Clone(s.data)}
}
