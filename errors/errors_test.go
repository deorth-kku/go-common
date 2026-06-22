package cerrors

import (
	"errors"
	"fmt"
	"testing"
)

type interror int

func (i interror) Error() string {
	return fmt.Sprintf("%d", i)
}

func TestUnwraps(t *testing.T) {
	err := errors.New("start")
	for i := range interror(10) {
		err = fmt.Errorf("this is %w, %w", err, error(i))
	}
	for i := range Unwraps(err) {
		fmt.Println(i.Error())
	}
	fmt.Println(err.Error())
}
