package validator

import "github.com/go-playground/validator/v10"

type RequestValidator interface {
	Validate(i interface{}) []string
}

type requestValidator struct {
	validator *validator.Validate
}

func NewRequestValidator() RequestValidator {
	return &requestValidator{
		validator: validator.New(),
	}
}

func (r *requestValidator) Validate(i interface{}) []string {
	if err := r.validator.Struct(i); err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, err.Field()+" is invalid: "+err.Tag())
		}

		return errors
	}

	return nil
}
