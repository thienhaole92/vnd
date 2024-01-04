package validator

import "github.com/go-playground/validator/v10"

type RestValidator struct {
	Validator *validator.Validate
}

func DefaultRestValidator() *RestValidator {
	r := &RestValidator{Validator: validator.New()}
	return r
}

func (v *RestValidator) Validate(i interface{}) error {
	if err := v.Validator.Struct(i); err != nil {
		return err
	}
	return nil
}
