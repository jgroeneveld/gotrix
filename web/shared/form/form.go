package form

import (
	"fmt"
	"github.com/jgroeneveld/bookie2/web/shared/httperr"
	"net/http"
	"strconv"
)

func New(r *http.Request) (*Form, error) {
	err := r.ParseForm()
	if err != nil {
		return nil, err
	}
	return &Form{
		Values: r.Form,
	}, nil
}

// Form wraps accessing of form data into accesors.
// Accessors can be Req* or Get*. Req* accesors require a attribute to be set and will otherwise
// store an error if not given. This will not protect against empty strings, 0 etc.
// Get* accessors give nil if not given.
type Form struct {
	Values      map[string][]string
	fieldErrors map[string][]string
}

func (f *Form) ReqString(field string) string {
	v, ok := f.Values[field]
	if !ok || len(v) == 0 {
		f.addErrorMissingParameter(field)
		return ""
	}

	return v[0]
}

func (f *Form) ReqInt(field string) int {
	v, ok := f.Values[field]
	if !ok || len(v) == 0 {
		f.addErrorMissingParameter(field)
		return 0
	}

	i, err := strconv.Atoi(v[0])
	if err != nil {
		f.addErrorParsing(field, "int")
	}
	return i
}

func (f *Form) ReqBool(field string) bool {
	v, ok := f.Values[field]
	if !ok || len(v) == 0 {
		f.addErrorMissingParameter(field)
		return false
	}

	b, err := strconv.ParseBool(v[0])
	if err != nil {
		f.addErrorParsing(field, "bool")
	}
	return b
}

func (f *Form) GetString(field string) *string {
	v, ok := f.Values[field]
	if !ok || len(v) == 0 {
		return nil
	}

	str := v[0]
	return &str
}

func (f *Form) GetInt(field string) *int {
	v, ok := f.Values[field]
	if !ok || len(v) == 0 {
		return nil
	}

	i, err := strconv.Atoi(v[0])
	if err != nil {
		f.addErrorParsing(field, "int")
		return nil
	}
	return &i
}

func (f *Form) GetBool(field string) *bool {
	v, ok := f.Values[field]
	if !ok || len(v) == 0 {
		return nil
	}

	b, err := strconv.ParseBool(v[0])
	if err != nil {
		f.addErrorParsing(field, "bool")
		return nil
	}
	return &b
}

func (f *Form) Err() error {
	if len(f.fieldErrors) == 0 {
		return nil
	}
	return httperr.Validation(f.fieldErrors)
}

func (f *Form) addErrorMissingParameter(field string) {
	f.addError(field, "is missing")
}

func (f *Form) addErrorParsing(field string, parseType string) {
	f.addError(field, fmt.Sprintf("can not be parsed to %s", parseType))
}

func (f *Form) addError(field, msg string) {
	if f.fieldErrors == nil {
		f.fieldErrors = map[string][]string{}
	}
	f.fieldErrors[field] = append(f.fieldErrors[field], msg)
}
