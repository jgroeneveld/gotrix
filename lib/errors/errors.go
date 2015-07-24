package errors

import (
	"fmt"
	"github.com/go-errors/errors"
)

func Wrap(err error) error {
	if err == nil {
		return err
	}
	return errors.Wrap(err, 1)
}

func New(format string, args ...interface{}) error {
	return errors.Wrap(fmt.Errorf(format, args...), 1)
}

func Is(e error, original error) bool {
	return errors.Is(e, original)
}

func GetOriginalAndStack(err error) (original error, stack string) {
	if casted, ok := err.(*errors.Error); ok {
		return casted.Err, casted.ErrorStack()
	}
	return err, ""
}

func ErrorWithStack(err error) string {
	if casted, ok := err.(interface {
		ErrorStack() string
	}); ok {
		return casted.ErrorStack()
	}
	return err.Error()
}
