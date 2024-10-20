package common

import (
	"context"
	"time"
)

type float interface {
	~float32 | ~float64
}

func FloatDuration[T float](f T) time.Duration {
	return time.Duration(T(time.Second) * f)
}
func Duration[T float](t time.Duration) T {
	return T(t.Seconds())
}

func TimeoutContext[T float](f T) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), FloatDuration(f))
}
