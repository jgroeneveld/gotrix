package httperr

import "github.com/jgroeneveld/bookie2/app/apperr"

// Convert converts application level errors into http errors
func Convert(err error) *Error {
	if err == nil {
		return nil
	}

	if httpErr, ok := err.(*Error); ok {
		// it is a httperror already
		return httpErr
	}

	if e, ok := err.(*apperr.Error); ok {
		if ve, ok := e.Err.(*apperr.ValidationError); ok {
			return Validation(ve.FieldErrors)
		}

		switch e.Err {
		case apperr.RecordNotFoundErr:
			return NotFound()
		}
	}

	return InternalServerError(err.Error())
}
