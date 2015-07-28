package httperr

import (
	apperrors "gotrix/app/errors"
	"gotrix/lib/errors"
)

// Convert converts application level errors into http errors
func Convert(err error) *Error {
	if err == nil {
		return nil
	}

	err, stack := errors.GetOriginalAndStack(err)

	if httpErr, ok := err.(*Error); ok {
		// it is a httperror already
		return httpErr
	}

	if ae, ok := err.(*apperrors.Error); ok {
		if ve, ok := ae.IsValidationError(); ok {
			return Validation(ve.FieldErrors)
		}

		if ok := ae.IsRecordNotFoundError(); ok {
			return NotFound()
		}
	}

	return InternalServerError(err.Error(), stack)
}
