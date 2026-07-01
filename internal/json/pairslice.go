package ijson

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"reflect"
	"slices"
)

type Pair[K, V any] = struct {
	Key   K
	Value V
}

func UnmarshalPairSlice[S ~[]P, P ~Pair[K, V], K, V any](dec *jsontext.Decoder, ps *S) error {
	tok, err := dec.ReadToken()
	if err != nil {
		return err
	}
	switch tok.Kind() {
	case '{':
	case 'n':
		*ps = nil
		return nil
	default:
		return &json.SemanticError{
			ByteOffset:  dec.InputOffset(),
			JSONPointer: dec.StackPointer(),
			JSONKind:    tok.Kind(),
			GoType:      reflect.TypeFor[S](),
		}
	}
	result := make([]P, 0)
	var line Pair[K, V]
	for dec.PeekKind() != '}' {
		line = Pair[K, V]{}
		err = json.UnmarshalDecode(dec, &line.Key)
		if err != nil {
			return err
		}
		err = json.UnmarshalDecode(dec, &line.Value)
		if err != nil {
			return err
		}
		result = append(result, line)
	}
	_, err = dec.ReadToken()
	if err != nil {
		return err
	}
	result = slices.Clip(result)
	*ps = result
	return nil
}

func MarshalPairSlice[S ~[]P, P ~Pair[K, V], K, V any](enc *jsontext.Encoder, ps S) error {
	if ps == nil {
		return enc.WriteToken(jsontext.Null)
	}
	err := enc.WriteToken(jsontext.BeginObject)
	if err != nil {
		return err
	}
	for _, p0 := range ps {
		p := Pair[K, V](p0)
		err = json.MarshalEncode(enc, p.Key)
		if err != nil {
			return err
		}
		err = json.MarshalEncode(enc, p.Value)
		if err != nil {
			return err
		}
	}
	return enc.WriteToken(jsontext.EndObject)
}
