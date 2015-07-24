package validate

import "github.com/jgroeneveld/gotrix/app/errors"

type Validator struct {
	FieldErrors map[string][]string
}

func NewValidator() *Validator {
	return &Validator{
		FieldErrors: map[string][]string{},
	}
}

func (v *Validator) Add(field string, msg string) {
	if v.FieldErrors == nil {
		v.FieldErrors = map[string][]string{}
	}
	v.FieldErrors[field] = append(v.FieldErrors[field], msg)
}

func (v *Validator) Err() error {
	if len(v.FieldErrors) > 0 {
		return errors.Validation(v.FieldErrors)
	}
	return nil
}
