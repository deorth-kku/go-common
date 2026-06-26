package cweak

import (
	"runtime"
	"strconv"
	"weak"
)

const (
	// CancelOp signals to Compute to not do anything as a result
	// of executing the lambda. If the entry was not present in
	// the map, nothing happens, and if it was present, the
	// returned value is ignored.
	CancelOp = iota
	// UpdateOp signals to Compute to update the entry to the
	// value returned by the lambda, creating it if necessary.
	UpdateOp
	// DeleteOp signals to Compute to always delete the entry
	// from the map.
	DeleteOp
)

type mapImp[K comparable, V any] map[K]V

func (m mapImp[K, V]) Compute(key K, valueFn func(V, bool) (V, int)) (V, bool) {
	new, op := valueFn(m.Load(key))
	switch op {
	case CancelOp:
		return new, false
	case UpdateOp:
		m[key] = new
		return new, true
	case DeleteOp:
		delete(m, key)
		return new, false
	default:
		panic("invalid op: " + strconv.Itoa(op))
	}
}

func (m mapImp[K, V]) Load(key K) (V, bool) {
	v, ok := m[key]
	return v, ok
}

func (m mapImp[K, V]) Delete(key K) {
	delete(m, key)
}

func newmap[K comparable, V any]() mapImp[K, V] {
	return make(mapImp[K, V])
}

// standard weak map, not concurrent-safe
func NewMap[K comparable, V any]() Map[K, V] {
	return NewCustomMap(newmap[K, Entry[V]])
}

// can be used with xsync.Map to get a concurrent-safe weak map
func NewCustomMap[SM mapInf[K, Entry[V], OP], K comparable, V any, OP ~int](f func() SM) Map[K, V] {
	return weakMap[K, V, OP, SM]{m: f()}
}

type Map[K comparable, V any] interface {
	Store(K, *V)
	Load(K) *V
}

type mapInf[K comparable, V any, OP ~int] interface {
	Compute(K, func(V, bool) (V, OP)) (V, bool)
	Load(K) (V, bool)
	Delete(K)
}

type weakMap[K comparable, V any, OP ~int, SM mapInf[K, Entry[V], OP]] struct {
	m SM
}

type Entry[V any] struct {
	p weak.Pointer[V]
	c runtime.Cleanup
}

func (m weakMap[K, V, OP, SM]) Store(k K, p *V) {
	if p == nil {
		return
	}
	m.m.Compute(k, func(oldValue Entry[V], loaded bool) (newValue Entry[V], op OP) {
		if loaded {
			oldValue.c.Stop()
		}
		return Entry[V]{
			p: weak.Make(p),
			c: runtime.AddCleanup(p, m.m.Delete, k),
		}, OP(UpdateOp)
	})
}

func (m weakMap[K, V, OP, SM]) Load(k K) *V {
	v, ok := m.m.Load(k)
	if !ok {
		return nil
	}
	return v.p.Value()
}
