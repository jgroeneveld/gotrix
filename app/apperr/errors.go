// apperr contains application level errors. They will be converted by the presentation layer (e.g. web)
package apperr

import (
	"errors"
	"fmt"
	"strings"
)

func Validation(fieldErrors map[string][]string) error {
	return &Error{
		Err: &ValidationError{FieldErrors: fieldErrors},
	}
}

func RecordNotFound() error {
	return &Error{
		Err: RecordNotFoundErr,
	}
}

var RecordNotFoundErr = errors.New("RecordNotFound")

type ValidationError struct {
	FieldErrors map[string][]string
}

func (e *ValidationError) Error() string {
	var msgs []string
	for k, v := range e.FieldErrors {
		msgs = append(msgs, fmt.Sprintf("%s: %s", k, strings.Join(v, ",")))
	}

	return fmt.Sprintf("ValidationError: %s", strings.Join(msgs, ", "))
}

type Error struct {
	Err error
}

func (e *Error) Error() string {
	return "apperror." + e.Err.Error()
}
