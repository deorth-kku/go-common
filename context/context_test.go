package ccontext

import (
	"context"
	"testing"

	"github.com/deorth-kku/go-common/args"
	ctest "github.com/deorth-kku/go-common/test"
)

func TestSplitContext(t *testing.T) {
	a := t.Context()
	b := args.Drop1(context.WithCancel(a))
	merged, _ := MergeContext(a, b)
	a2, b2 := SplitContext(merged)

	ctest.Equal(t, a, a2)
	ctest.Equal(t, b, b2)
}
