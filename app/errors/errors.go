// errors contains application level errors. They will be converted by the presentation layer (e.g. web)
package errors

import (
	"fmt"
	"gotrix/lib/errors"
	"strings"
)

func Validation(fieldErrors map[string][]string) error {
	return errors.Wrap(&Error{
		Err: &validationError{FieldErrors: fieldErrors},
	})
}

func RecordNotFound() error {
	return errors.Wrap(&Error{
		Err: recordNotFoundErr,
	})
}

// Use own type to identify application level errors, allow error vars as errors and custom data attachments.
type Error struct {
	Err error
}

func (e *Error) Error() string {
	return e.Err.Error()
}

func (err *Error) IsValidationError() (*validationError, bool) {
	ve, ok := err.Err.(*validationError)
	return ve, ok
}

func (err *Error) IsRecordNotFoundError() bool {
	return err.Err == recordNotFoundErr
}

var recordNotFoundErr = fmt.Errorf("RecordNotFound")

type validationError struct {
	FieldErrors map[string][]string
}

func (e *validationError) Error() string {
	var msgs []string
	for k, v := range e.FieldErrors {
		msgs = append(msgs, fmt.Sprintf("%s: %s", k, strings.Join(v, ",")))
	}

	return fmt.Sprintf("ValidationError: %s", strings.Join(msgs, ", "))
}
