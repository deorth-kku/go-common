package ccontext

import (
	"context"
	"reflect"
	"time"
)

type mergedContext struct {
	left, right context.Context
}

func (m *mergedContext) Deadline() (time.Time, bool) {
	a, ok := m.left.Deadline()
	if !ok {
		return m.right.Deadline()
	}
	b, ok := m.right.Deadline()
	if !ok {
		return a, true
	}
	if a.After(b) {
		return b, true
	}
	return a, true
}

func (m *mergedContext) Value(key any) any {
	v := m.left.Value(key)
	if v == nil {
		return m.right.Value(key)
	}
	return v
}

func (m *mergedContext) Done() <-chan struct{} {
	return m.right.Done()
}

func (m *mergedContext) Err() error {
	aerr := m.left.Err()
	if aerr == nil {
		return m.right.Err()
	}
	return aerr
}

func MergeContext(left, right context.Context) (context.Context, context.CancelFunc) {
	switch left {
	case context.Background(), context.TODO(), nil:
		return context.WithCancel(right)
	}
	switch right {
	case context.Background(), context.TODO(), nil:
		return context.WithCancel(left)
	}
	rightalt, cancel := context.WithCancel(right)
	stop := context.AfterFunc(left, cancel)
	return &mergedContext{left, rightalt}, func() {
		stop()
		cancel()
	}
}

func SplitContext(merged context.Context) (left, right context.Context) {
	mer, ok := merged.(*mergedContext)
	if !ok {
		return
	}
	left = mer.left
	rv := reflect.ValueOf(mer.right).Elem().Field(0)
	right, ok = reflect.TypeAssert[context.Context](rv)
	if !ok {
		panic("context package changed")
	}
	return
}
