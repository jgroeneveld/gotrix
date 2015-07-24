package httperr

import (
	"encoding/json"
	"fmt"
)

const (
	StatusInternalServerError = 500
	StatusBadRequest          = 400
	StatusMissingParameter    = 400
	StatusUnauthorized        = 401
	StatusNotFound            = 404
	StatusMethodNotAllowed    = 405
	StatusValidationError     = 422
	StatusFailedDependency    = 424
)

func InternalServerError(msg string, stacktrace string) *Error {
	return &Error{
		Status:     StatusInternalServerError,
		Type:       "internal_server_error",
		Message:    "Internal Server Error: " + msg,
		Stacktrace: stacktrace,
	}
}

func BadRequest(msg string) error {
	return &Error{
		Status:  StatusBadRequest,
		Type:    "bad_request",
		Message: "Bad request: " + msg,
	}
}

func MissingParameter(param string) error {
	return &Error{
		Status:  StatusMissingParameter,
		Type:    "missing_parameter",
		Message: fmt.Sprintf("'%s' must be given", param),
	}
}

func Unauthorized() error {
	return &Error{
		Status:  StatusUnauthorized,
		Type:    "unauthorized",
		Message: "You are not authorized",
	}
}

func NotFound() *Error {
	return &Error{
		Status:  StatusNotFound,
		Type:    "not_found",
		Message: "The requested resource could not be found",
	}
}

func MethodNotAllowed() error {
	return &Error{
		Status:  StatusMethodNotAllowed,
		Type:    "method_not_allowed",
		Message: "Method not allowed",
	}
}

// TODO make sure we can return *Error instead of error
func Validation(fieldErrors map[string][]string) *Error {
	return &Error{
		Status:  StatusValidationError,
		Type:    "validation",
		Message: "Validation failure",
		Errors:  fieldErrors,
	}
}

func FailedDependency(msg string) error {
	return &Error{
		Status:  StatusFailedDependency,
		Type:    "failed_dependency",
		Message: msg,
	}
}

type Error struct {
	Status     int                 `json:"status"`
	Type       string              `json:"type"`
	Message    string              `json:"message"`
	Errors     map[string][]string `json:"errors,omitempty"`
	Stacktrace string              `json:"-"`
}

func (err Error) Error() string {
	suffix := ""
	if len(err.Errors) > 0 {
		b, _ := json.Marshal(err.Errors)
		suffix = fmt.Sprintf(" %s", string(b))
	}
	return fmt.Sprintf("%s(%d): %s%s", err.Type, err.Status, err.Message, suffix)
}

// MarshalJSON wraps the error in {"error":...}
func (err Error) MarshalJSON() ([]byte, error) {
	type tempErr Error

	x := struct {
		Error tempErr `json:"error"`
	}{tempErr(err)}

	return json.Marshal(x)
}

func (err *Error) UnmarshalJSON(b []byte) error {
	type tempErr Error

	x := struct {
		Error tempErr `json:"error"`
	}{}

	jsonErr := json.Unmarshal(b, &x)
	if jsonErr != nil {
		return jsonErr
	}

	*err = Error(x.Error)
	return nil
}
