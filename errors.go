package common

import (
	"errors"
	"slices"
)

type Unwraper interface {
	Unwrap() error
}

type MultiUnwraper interface {
	Unwrap() []error
}

func Unwraps(err error) Seq[error] {
	return func(yield Yield[error]) {
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

func MergeError(msg string, errs ...error) error {
	errs = slices.Insert(errs, 0, error(ErrorString(msg)))
	return errors.Join(errs...)
}

func Errors(msg string, errs ...error) error {
	if len(errs) == 0 {
		return nil
	}
	return MergeError(msg, errs...)
}

type ErrorString string

func (e ErrorString) Error() string {
	return string(e)
}
