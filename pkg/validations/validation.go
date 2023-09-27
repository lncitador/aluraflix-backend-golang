package validations

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

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
			return new(bool), []string{field.Error()}
		}

		return new(bool), nil
	}

	return nil, nil
}
