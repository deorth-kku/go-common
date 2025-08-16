package common

import (
	"cmp"
	"maps"
	"slices"
)

type (
	Yield[T any]     = func(T) bool
	Yield2[K, V any] = func(K, V) bool
	Seq[T any]       = func(Yield[T])
	Seq2[K, V any]   = func(Yield2[K, V])
	Ranger[T any]    interface {
		Range(Yield[T])
	}
	Ranger2[K, V any] = interface {
		Range(Yield2[K, V])
	}
)

type Pair[K any, V any] struct {
	Key   K
	Value V
}

func NewPair[K any, V any](Key K, Value V) Pair[K, V] {
	return Pair[K, V]{Key: Key, Value: Value}
}

func PairSliceCollect[K any, V any](iter Seq2[K, V], hint ...int) (pairs PairSlice[K, V]) {
	if len(hint) != 0 {
		pairs = make(PairSlice[K, V], 0, hint[0])
	}
	for k, v := range iter {
		pairs = append(pairs, NewPair(k, v))
	}
	return
}
func PairSliceFromMap[K comparable, V any, M ~map[K]V](m M) (pairs PairSlice[K, V]) {
	return PairSliceCollect(maps.All(m), len(m))
}

type PairSlice[K any, V any] []Pair[K, V]

func (ps PairSlice[K, V]) Keys(yield Yield[K]) {
	slices.Values(ps)(func(line Pair[K, V]) bool {
		return yield(line.Key)
	})
}

func (ps PairSlice[K, V]) SliceKeys() []K {
	return SliceCollect(ps.Keys, len(ps))
}

func (ps PairSlice[K, V]) Values(yield Yield[V]) {
	slices.Values(ps)(func(line Pair[K, V]) bool {
		return yield(line.Value)
	})
}

func (ps PairSlice[K, V]) SliceValues() []V {
	return SliceCollect(ps.Values, len(ps))
}

func (ps PairSlice[K, V]) Range(yield Yield2[K, V]) {
	slices.Values(ps)(func(line Pair[K, V]) bool {
		return yield(line.Key, line.Value)
	})
}

func (ps PairSlice[K, V]) Backward(yield Yield2[K, V]) {
	slices.Backward(ps)(func(_ int, line Pair[K, V]) bool {
		return yield(line.Key, line.Value)
	})
}

func (ps PairSlice[K, V]) search(key K, equal func(a, b K) bool) (V, bool) {
	for _, line := range ps {
		if equal(line.Key, key) {
			return line.Value, true
		}
	}
	var v V
	return v, false
}

func Search[K comparable, V any](ps PairSlice[K, V], key K) (V, bool) {
	return ps.search(key, Equal)
}

func SearchEqual[K CanEqual[K], V any](ps PairSlice[K, V], key K) (V, bool) {
	return ps.search(key, EqualT)
}

func SearchT[K CanCompare[K], V any](ps PairSlice[K, V], key K) (V, bool) {
	return ps.search(key, func(a, b K) bool {
		return a.Compare(b) == 0
	})
}

func Sort[K cmp.Ordered, V any](ps PairSlice[K, V]) {
	slices.SortFunc(ps, func(a, b Pair[K, V]) int {
		return cmp.Compare(a.Key, b.Key)
	})
}

func SortT[K CanCompare[K], V any](ps PairSlice[K, V]) {
	slices.SortFunc(ps, func(a, b Pair[K, V]) int {
		return CompareT(a.Key, b.Key)
	})
}

type (
	cmpFunc[K any]     = func(x K, y K) int
	computeFunc[V any] = func(oldValue V, loaded bool) (newValue V, delete bool)
)

func (ps PairSlice[K, V]) insert(key K, value V, cmp cmpFunc[K]) PairSlice[K, V] {
	id, ok := slices.BinarySearchFunc(ps, NewPair(key, value), func(a, b Pair[K, V]) int {
		return cmp(a.Key, b.Key)
	})
	if ok {
		ps[id].Value = value
	} else {
		ps = slices.Insert(ps, id, NewPair(key, value))
	}
	return ps
}

func Insert[K cmp.Ordered, V any](ps PairSlice[K, V], key K, value V) PairSlice[K, V] {
	return ps.insert(key, value, cmp.Compare)
}

func InsertT[K CanCompare[K], V any](ps PairSlice[K, V], key K, value V) PairSlice[K, V] {
	return ps.insert(key, value, CompareT)
}

func (ps PairSlice[K, V]) compute(key K, valueFn computeFunc[V], cmp cmpFunc[K]) (PairSlice[K, V], V, bool) {
	p := Pair[K, V]{Key: key}
	id, ok := slices.BinarySearchFunc(ps, p, func(a, b Pair[K, V]) int {
		return cmp(a.Key, b.Key)
	})
	if ok {
		p.Value, ok = valueFn(ps[id].Value, ok)
		if ok {
			ps = slices.Delete(ps, id, id+1)
		} else {
			ps[id].Value = p.Value
		}
	} else {
		p.Value, ok = valueFn(p.Value, ok)
		if !ok {
			ps = slices.Insert(ps, id, p)
		}
	}
	return ps, p.Value, !ok
}

func Compute[K cmp.Ordered, V any](ps PairSlice[K, V], key K, valueFn computeFunc[V]) (PairSlice[K, V], V, bool) {
	return ps.compute(key, valueFn, cmp.Compare)
}

func ComputeT[K CanCompare[K], V any](ps PairSlice[K, V], key K, valueFn computeFunc[V]) (PairSlice[K, V], V, bool) {
	return ps.compute(key, valueFn, CompareT)
}

func (ps PairSlice[K, V]) binarySearch(key K, cmp cmpFunc[K]) (V, bool) {
	p := Pair[K, V]{Key: key}
	id, ok := slices.BinarySearchFunc(ps, p, func(a, b Pair[K, V]) int {
		return cmp(a.Key, b.Key)
	})
	if ok {
		return ps[id].Value, true
	}
	return p.Value, false
}

func BinarySearch[K cmp.Ordered, V any](ps PairSlice[K, V], key K) (V, bool) {
	return ps.binarySearch(key, cmp.Compare)
}

func BinarySearchT[K CanCompare[K], V any](ps PairSlice[K, V], key K) (V, bool) {
	return ps.binarySearch(key, CompareT)
}

func (ps PairSlice[K, V]) delete(key K, cmp cmpFunc[K]) PairSlice[K, V] {
	id, ok := slices.BinarySearchFunc(ps, Pair[K, V]{Key: key}, func(a, b Pair[K, V]) int {
		return cmp(a.Key, b.Key)
	})
	if ok {
		ps = slices.Delete(ps, id, id+1)
	}
	return ps
}

func Delete[K cmp.Ordered, V any](ps PairSlice[K, V], key K) PairSlice[K, V] {
	return ps.delete(key, cmp.Compare)
}

func DeleteT[K CanCompare[K], V any](ps PairSlice[K, V], key K) PairSlice[K, V] {
	return ps.delete(key, CompareT)
}

type psdata[K any, V any] = PairSlice[K, V]

type BSMap[K cmp.Ordered, V any] struct {
	psdata[K, V]
}

func NewBSMap[K cmp.Ordered, V any](from ...Pair[K, V]) *BSMap[K, V] {
	Sort(from)
	return &BSMap[K, V]{from}
}

func (bs BSMap[K, V]) Load(key K) (V, bool) {
	return bs.binarySearch(key, cmp.Compare)
}

func (bs *BSMap[K, V]) Store(key K, value V) {
	bs.psdata = bs.insert(key, value, cmp.Compare)
}

func (bs *BSMap[K, V]) Delete(key K) {
	bs.psdata = bs.delete(key, cmp.Compare)
}

func (bs *BSMap[K, V]) Compute(key K, valueFn computeFunc[V]) (actual V, ok bool) {
	bs.psdata, actual, ok = bs.compute(key, valueFn, cmp.Compare)
	return
}

func (bs BSMap[K, V]) Size() int {
	return len(bs.psdata)
}

func NewBSMapT[K CanCompare[K], V any](from ...Pair[K, V]) *BSMapT[K, V] {
	SortT(from)
	return &BSMapT[K, V]{from}
}

type BSMapT[K CanCompare[K], V any] struct {
	psdata[K, V]
}

func (bs BSMapT[K, V]) Load(key K) (V, bool) {
	return bs.binarySearch(key, CompareT)
}

func (bs *BSMapT[K, V]) Store(key K, value V) {
	bs.psdata = bs.insert(key, value, CompareT)
}

func (bs *BSMapT[K, V]) Delete(key K) {
	bs.psdata = bs.delete(key, CompareT)
}

func (bs *BSMapT[K, V]) Compute(key K, valueFn computeFunc[V]) (actual V, ok bool) {
	bs.psdata, actual, ok = bs.compute(key, valueFn, CompareT)
	return
}

func (bs BSMapT[K, V]) Size() int {
	return len(bs.psdata)
}

func EmptyRange[T any](Yield[T])             {}
func EmptyRange2[K any, V any](Yield2[K, V]) {}

func SafeRange[IT ~Seq[T], T any](fs ...IT) IT {
	return func(yield Yield[T]) {
		for i, f := range fs {
			if f == nil {
				continue
			}
			if i == len(fs)-1 {
				f(yield)
				return
			}
			for v := range f {
				if !yield(v) {
					return
				}
			}
		}
	}
}

func SafeRange2[IT ~Seq2[K, V], K, V any](fs ...IT) IT {
	return func(yield Yield2[K, V]) {
		for i, f := range fs {
			if f == nil {
				continue
			}
			if i == len(fs)-1 {
				f(yield)
				return
			}
			for k, v := range f {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}

type PairChan[K any, V any] chan Pair[K, V]

func (pc PairChan[K, V]) Range(yield Yield2[K, V]) {
	for pair := range pc {
		if !yield(pair.Key, pair.Value) {
			return
		}
	}
}

func Seq2K[K any, V any](it Seq2[K, V]) Seq[K] {
	return func(yield Yield[K]) {
		it(func(k K, _ V) bool {
			return yield(k)
		})
	}
}

func Seq2V[K any, V any](it Seq2[K, V]) Seq[V] {
	return func(yield Yield[V]) {
		it(func(_ K, v V) bool {
			return yield(v)
		})
	}
}

func Filter[T any](seq Seq[T], filter Yield[T]) Seq[T] {
	return func(yield Yield[T]) {
		seq(func(v T) bool {
			if filter(v) {
				return yield(v)
			}
			return true
		})
	}
}
func Filter2[K any, V any](seq Seq2[K, V], filter Yield2[K, V]) Seq2[K, V] {
	return func(yield Yield2[K, V]) {
		seq(func(k K, v V) bool {
			if filter(k, v) {
				return yield(k, v)
			}
			return true
		})
	}
}
