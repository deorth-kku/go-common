package common

import (
	"context"
	"time"
)

type Float interface {
	~float32 | ~float64
}

func FloatDuration[T Float](f T) time.Duration {
	return time.Duration(T(time.Second) * f)
}
func Duration[T Float](t time.Duration) T {
	return T(t.Seconds())
}

func TimeoutContext[T Float](f T) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), FloatDuration(f))
}
