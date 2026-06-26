package cweak

import (
	"fmt"
	"reflect"
	"weak"
)

type weakInterface[I, T any] struct {
	p weak.Pointer[T]
}

func (w weakInterface[I, T]) Value() I {
	p := w.p.Value()
	if p == nil {
		var i I
		return i
	}
	return any(p).(I)
}

func (w weakInterface[I, T]) toany() any {
	p := w.p.Value()
	if p == nil {
		return nil
	}
	return any(p)
}

func NewInterface[I, T any](ptr *T) Interface[I] {
	if _, ok := any(ptr).(I); !ok {
		ptrTy := reflect.TypeFor[*T]()
		infTy := reflect.TypeFor[I]()
		panic(fmt.Sprintf("%s does not implement %s", ptrTy.Name(), infTy.Name()))
	}
	if ptr == nil {
		return nil
	}
	wp := weak.Make(ptr)
	return weakInterface[I, T]{wp}
}

type Interface[T any] interface {
	toany() any // ensure only [weakInterface] can implement this
	Value() T
}

func TypeAssert[T, U any](w Interface[U]) (T, bool) {
	v, ok := w.toany().(T)
	return v, ok
}
