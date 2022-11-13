package errorutil

import (
	"fmt"

	"github.com/pkg/errors"
)

type customError struct {
	code Code
	err  error
}

func (e customError) Error() string {
	return fmt.Sprintf("%s", e.err)
}

func ErrorCode(err error) Code {
	if err == nil {
		return OK
	}

	e, ok := err.(customError)
	if ok {
		return e.code
	}

	return Unknown
}

func Errorf(code Code, format string, args ...interface{}) error {
	if code == OK {
		return nil
	}

	err := customError{
		code: code,
		err:  errors.Errorf(format, args...),
	}

	err.PrintCodeAndMsg()

	return err
}

func (e customError) PrintCodeAndMsg() {
	fmt.Printf("Code: %d, Msg: %s\n", e.code.Number(), e.Error())
}
