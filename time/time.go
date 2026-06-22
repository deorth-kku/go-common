package ctime

import (
	"context"
	"time"

	cmath "github.com/deorth-kku/go-common/math"
)

func ToDuration[T cmath.Float](f T) time.Duration {
	return time.Duration(T(time.Second) * f)
}
func DurationTo[T cmath.Float](t time.Duration) T {
	return T(t.Seconds())
}

func TimeoutContext[T cmath.Float](f T) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), ToDuration(f))
}
