package validations

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

// ErrosAsValidationByField checks if an error is a validation error related to a specific field.
//
// This function accepts two parameters:
// - err: The error to be checked.
// - by: A string specifying the field to look for validation errors.
//
// Returns:
//   - A pointer to a boolean that is true if the error is a field validation error
//     related to the field specified by 'by'. Returns false if the error is not a validation error.
//   - A string containing the error message associated with the field validation error, if applicable.
//
// The function checks if the error is a validation error using the Go "validator" package.
// If the error is a list of validation errors (ValidationErrors), the function checks if any of them
// are associated with the field specified by 'by' and returns the first error found, along with the
// error message. If the error is a single field validation error (FieldError), the function checks
// if it is associated with the field specified by 'by' and returns the error along with the error message.
//
// If the error is not a validation error, the function returns false for the boolean pointer and an empty string.
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
