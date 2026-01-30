package common

import (
	"strings"
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

var _ MultiUnwraper = mergeError{}

type mergeError struct {
	msg  string
	errs []error
}

func (m mergeError) Error() string {
	if len(m.errs) == 0 {
		return m.msg
	}
	bld := new(strings.Builder)
	if len(m.msg) != 0 {
		bld.WriteString(m.msg)
		bld.WriteString(": ")
	}
	bld.WriteString(m.errs[0].Error())
	for _, err := range m.errs[1:] {
		bld.WriteString(",")
		bld.WriteString(err.Error())
	}
	return bld.String()
}

func (m mergeError) Unwrap() []error {
	if len(m.errs) == 0 {
		return nil
	}
	return m.errs
}

func MergeError(msg string, errs ...error) mergeError {
	return mergeError{msg, errs}
}

func Errors(msg string, errs ...error) error {
	if len(errs) == 0 {
		return nil
	}
	return mergeError{msg, errs}
}

type ErrorString string

func (e ErrorString) Error() string {
	return string(e)
}
