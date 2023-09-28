package validations

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

// ErrosAsValidationByField returns a boolean indicating if the error is a validation error and the error message
func ErrosAsValidationByField(err error, by string) (*bool, string) {
	errValidation := true
	if err != nil {
		var fields validator.ValidationErrors
		var field validator.FieldError
		if errors.As(err, &fields) {
			for _, field := range fields {
				if field.Field() == by {
					return &errValidation, field.Error()
				}
			}
		} else if errors.As(err, &field) {
			return &errValidation, field.Error()
		}

		return new(bool), ""
	}

	return nil, ""
}

// GetErrorsByValidation returns a boolean indicating if the error is a validation error and the error message
// Deprecated: Use ErrosAsValidationByField instead
func GetErrorsByValidation(err error) (*bool, []string) {
	if err != nil {
		var fields validator.ValidationErrors
		var field validator.FieldError
		if errors.As(err, &fields) {
			var errs []string
			for _, field := range fields {
				errs = append(errs, field.Error())
			}
			return nil, errs
		} else if errors.As(err, &field) {
			return nil, []string{field.Error()}
		}

		return new(bool), nil
	}

	return nil, nil
}
