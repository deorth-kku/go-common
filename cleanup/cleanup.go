package cleanup

import "io"

func Closer[T io.Closer](c T) {
	c.Close()
}
