package cleanup

import "io"

func Closer[T io.Closer](c T) {
	c.Close()
}

func Func0[T ~func()](can T) {
	can()
}

func Func1[T ~func() R, R any](can T) {
	can()
}
