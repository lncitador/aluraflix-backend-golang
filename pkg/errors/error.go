package errors

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Error struct {
	Code        int       `json:"code"`
	Message     string    `json:"message"`
	Description *string   `json:"description,omitempty"`
	Fields      *[]string `json:"fields,omitempty"`
}

func (e Error) Error() string {
	return e.Message
}

func NewError(code int, message string, description string) *Error {
	return &Error{
		Code:        code,
		Message:     message,
		Description: &description,
		Fields:      nil,
	}
}

func NewErrorByInternal(err error) *Error {
	description := err.Error()
	return &Error{
		Code:        http.StatusInternalServerError,
		Message:     "Internal error",
		Description: &description,
		Fields:      nil,
	}
}

func NewErrorByValidation(err error) *Error {
	var fields validator.ValidationErrors
	var field validator.FieldError

	var errs []string
	if errors.As(err, &fields) {
		for _, field := range fields {
			errs = append(errs, field.Error())
		}
	} else if errors.As(err, &field) {
		errs = append(errs, field.Error())
	} else {
		description := err.Error()
		return &Error{
			Code:        http.StatusUnprocessableEntity,
			Message:     "Validation error",
			Description: &description,
			Fields:      nil,
		}
	}

	return &Error{
		Code:        http.StatusUnprocessableEntity,
		Message:     "Validation error",
		Description: nil,
		Fields:      &errs,
	}
}

func NewErrorByNotFound(err error) *Error {
	description := err.Error()
	return &Error{
		Code:        http.StatusNotFound,
		Message:     "Not found",
		Description: &description,
		Fields:      nil,
	}
}

func NewErrorByUnauthorized(err error) *Error {
	description := err.Error()
	return &Error{
		Code:        http.StatusUnauthorized,
		Message:     "Unauthorized",
		Description: &description,
		Fields:      nil,
	}
}

func NewErrorByForbidden(err error) *Error {
	description := err.Error()
	return &Error{
		Code:        http.StatusForbidden,
		Message:     "Forbidden",
		Description: &description,
		Fields:      nil,
	}
}

func NewErrorByBadRequest(err error) *Error {
	description := err.Error()
	return &Error{
		Code:        http.StatusBadRequest,
		Message:     "Bad request",
		Description: &description,
		Fields:      nil,
	}
}
