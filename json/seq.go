package cjson

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"iter"
	"reflect"
	"runtime"
	"slices"
	"sync"
	"weak"

	"github.com/deorth-kku/go-common/datatypes"
	ijson "github.com/deorth-kku/go-common/internal/json"
	citer "github.com/deorth-kku/go-common/iter"
)

// helper types to make [iter.Seq] and [iter.Seq2] works with encoding/json
type (
	SeqJson[T any]         iter.Seq[T]
	Seq2Json[K any, V any] iter.Seq2[K, V]
)

func (sj *SeqJson[T]) UnmarshalJSON(data []byte) error {
	return ijson.UnmarshalV1(data, sj)
}

func (sj *SeqJson[T]) UnmarshalJSONFrom(dec *jsontext.Decoder) error {
	var temp []T
	err := json.UnmarshalDecode(dec, &temp)
	if err != nil {
		return err
	}
	if temp == nil {
		return nil
	}
	temp = slices.Clip(temp)
	*sj = ToSeq(slices.Values(temp))
	store(*sj, &temp)
	return nil
}

func (sj SeqJson[T]) MarshalJSON() ([]byte, error) {
	return ijson.MarshalV1(sj)
}

func (sj SeqJson[T]) MarshalJSONTo(enc *jsontext.Encoder) error {
	if sj == nil {
		return enc.WriteToken(jsontext.Null)
	}
	err := enc.WriteToken(jsontext.BeginArray)
	if err != nil {
		return err
	}
	for line := range sj {
		err = json.MarshalEncode(enc, line)
		if err != nil {
			return err
		}
	}
	return enc.WriteToken(jsontext.EndArray)
}

func (sj *Seq2Json[K, V]) UnmarshalJSON(data []byte) error {
	return ijson.UnmarshalV1(data, sj)
}

func (sj *Seq2Json[K, V]) UnmarshalJSONFrom(dec *jsontext.Decoder) error {
	var result datatypes.PairSlice[K, V]
	err := ijson.UnmarshalPairSlice(dec, &result)
	if err != nil {
		return err
	}
	*sj = result.Range
	store2(*sj, &result)
	return nil
}

func (sj Seq2Json[K, V]) MarshalJSON() ([]byte, error) {
	return ijson.MarshalV1(sj)
}

func (sj Seq2Json[K, V]) MarshalJSONTo(enc *jsontext.Encoder) error {
	if sj == nil {
		return enc.WriteToken(jsontext.Null)
	}
	err := enc.WriteToken(jsontext.BeginObject)
	if err != nil {
		return err
	}
	for key, value := range sj {
		err = json.MarshalEncode(enc, key)
		if err != nil {
			return err
		}
		err = json.MarshalEncode(enc, value)
		if err != nil {
			return err
		}
	}
	return enc.WriteToken(jsontext.EndObject)
}

func ToSeq[T any](seq citer.Seq[T]) SeqJson[T] {
	return SeqJson[T](seq)
}

func ToSeq2[K, V any](seq citer.Seq2[K, V]) Seq2Json[K, V] {
	return Seq2Json[K, V](seq)
}

func SeqMarshalJSONTo[T any](seq citer.Seq[T], enc *jsontext.Encoder) error {
	return SeqJson[T](seq).MarshalJSONTo(enc)
}

func Seq2MarshalJSONTo[K, V any](seq citer.Seq2[K, V], enc *jsontext.Encoder) error {
	return Seq2Json[K, V](seq).MarshalJSONTo(enc)
}

func SeqMarshalJSON[T any](seq citer.Seq[T]) ([]byte, error) {
	return SeqJson[T](seq).MarshalJSON()
}

func Seq2MarshalJSON[K, V any](seq citer.Seq2[K, V]) ([]byte, error) {
	return Seq2Json[K, V](seq).MarshalJSON()
}

func ChanToSeq[T any](ch <-chan T) SeqJson[T] {
	return func(y citer.Yield[T]) {
		for it := range ch {
			if !y(it) {
				return
			}
		}
	}
}

func ChanToSeq2[K, V any](ch <-chan datatypes.Pair[K, V]) Seq2Json[K, V] {
	return datatypes.PairChan[K, V](ch).Range
}

var weakmap sync.Map

func toUintptr[T any](f T) uintptr {
	return reflect.ValueOf(f).Pointer()
}

func store[T any](key SeqJson[T], value *[]T) {
	ptrkey := toUintptr(key)
	p := weak.Make(value)
	runtime.AddCleanup(value, weakmap.Delete, any(ptrkey))
	weakmap.Store(ptrkey, p)
}

func get[T any](key SeqJson[T]) *[]T {
	p, ok := weakmap.Load(toUintptr(key))
	if !ok {
		return nil
	}
	return p.(weak.Pointer[[]T]).Value()
}

func store2[K, V any](key Seq2Json[K, V], value *datatypes.PairSlice[K, V]) {
	ptrkey := toUintptr(key)
	p := weak.Make(value)
	runtime.AddCleanup(value, weakmap.Delete, any(ptrkey))
	weakmap.Store(ptrkey, p)
}

func get2[K, V any](key Seq2Json[K, V]) *datatypes.PairSlice[K, V] {
	p, ok := weakmap.Load(toUintptr(key))
	if !ok {
		return nil
	}
	return p.(weak.Pointer[datatypes.PairSlice[K, V]]).Value()
}

func ToPairSlice[K, V any](it Seq2Json[K, V]) datatypes.PairSlice[K, V] {
	p := get2(it)
	if p != nil {
		return *p
	}
	return datatypes.PairSliceCollect(it)
}

func ToSlice[T any](it SeqJson[T]) []T {
	p := get(it)
	if p != nil {
		return *p
	}
	return slices.Collect(iter.Seq[T](it))
}
