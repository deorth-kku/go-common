package datatypes

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"maps"
	"reflect"

	cslices "github.com/deorth-kku/go-common/datatypes/slices"
	ijson "github.com/deorth-kku/go-common/internal/json"
	citer "github.com/deorth-kku/go-common/iter"
)

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

func (s Set[T]) Range(yield citer.Yield[T]) {
	if s.data == nil {
		return
	}
	maps.Keys(s.data)(yield)
}

func (s Set[T]) Slice() []T {
	if s.data == nil {
		return nil
	}
	return cslices.Collect(s.Range, len(s.data))
}

func (s Set[T]) Clone() Set[T] {
	return Set[T]{maps.Clone(s.data)}
}

func (s *Set[T]) UnmarshalJSON(data []byte) error {
	return ijson.UnmarshalV1(data, s)
}

func (s *Set[T]) UnmarshalJSONFrom(dec *jsontext.Decoder) error {
	tok, err := dec.ReadToken()
	if err != nil {
		return err
	}
	switch tok.Kind() {
	case '[':
	case 'n':
		return nil
	default:
		return &json.SemanticError{
			ByteOffset:  dec.InputOffset(),
			JSONPointer: dec.StackPointer(),
			JSONKind:    tok.Kind(),
			GoType:      reflect.TypeFor[Set[T]](),
		}
	}
	var temp Set[T]
	var line T
	for dec.PeekKind() != ']' {
		line = *new(T)
		err = json.UnmarshalDecode(dec, &line)
		if err != nil {
			return err
		}
		temp.Add(line)
	}
	_, err = dec.ReadToken()
	if err != nil {
		return err
	}
	s.data = temp.data
	return nil
}

func (s Set[T]) MarshalJSON() ([]byte, error) {
	return ijson.MarshalV1(s)
}

func (s Set[T]) MarshalJSONTo(enc *jsontext.Encoder) error {
	if s.data == nil {
		return enc.WriteToken(jsontext.Null)
	}
	err := enc.WriteToken(jsontext.BeginArray)
	if err != nil {
		return err
	}
	for line := range s.Range {
		err = json.MarshalEncode(enc, line)
		if err != nil {
			return err
		}
	}
	return enc.WriteToken(jsontext.EndArray)
}
