package cerrors

import (
	"errors"
	"slices"

	citer "github.com/deorth-kku/go-common/iter"
)

type Unwraper interface {
	Unwrap() error
}

type MultiUnwraper interface {
	Unwrap() []error
}

func Unwraps(err error) citer.Seq[error] {
	return func(yield citer.Yield[error]) {
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

func Merge(msg string, errs ...error) error {
	errs = slices.Insert(errs, 0, error(String(msg)))
	return errors.Join(errs...)
}

func Errors(msg string, errs ...error) error {
	if len(errs) == 0 {
		return nil
	}
	return Merge(msg, errs...)
}

type String string

func (e String) Error() string {
	return string(e)
}
