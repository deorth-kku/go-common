package common

import (
	"iter"
)

type Unwraper interface {
	Unwrap() error
}

type MultiUnwraper interface {
	Unwrap() []error
}

func Unwraps(err error) iter.Seq[error] {
	return func(yield func(error) bool) {
		switch t := err.(type) {
		case Unwraper:
			for e := range Unwraps(t.Unwrap()) {
				if !yield(e) {
					return
				}
			}
		case MultiUnwraper:
			for _, e0 := range t.Unwrap() {
				for e := range Unwraps(e0) {
					if !yield(e) {
						return
					}
				}
			}
		default:
			if !yield(err) {
				return
			}
		}
	}
}
