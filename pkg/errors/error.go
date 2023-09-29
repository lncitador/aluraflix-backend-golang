package errors

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Description string

type Error interface {
	Error() string
	Status() int
	GetDescription() string
}

type handle struct {
	Code        int       `json:"code"`
	Message     string    `json:"message"`
	Description *string   `json:"description,omitempty"`
	Fields      *[]string `json:"fields,omitempty"`
}

func (e handle) Error() string {
	return e.Message
}

func (e handle) Status() int {
	return e.Code
}

func (e handle) GetDescription() string {
	if e.Description != nil {
		return *e.Description
	}

	return ""
}

func NewError(code int, message string, description string) *handle {
	return &handle{
		Code:        code,
		Message:     message,
		Description: &description,
		Fields:      nil,
	}
}

func NewErrorByInternal(err error) *handle {
	description := err.Error()
	return &handle{
		Code:        http.StatusInternalServerError,
		Message:     "Internal error",
		Description: &description,
		Fields:      nil,
	}
}

func NewErrorByValidation(err error) *handle {
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
		return &handle{
			Code:        http.StatusUnprocessableEntity,
			Message:     "Validation error",
			Description: &description,
			Fields:      nil,
		}
	}

	return &handle{
		Code:        http.StatusUnprocessableEntity,
		Message:     "Validation error",
		Description: nil,
		Fields:      &errs,
	}
}

func NewErrorByNotFound(err error) *handle {
	description := err.Error()
	return &handle{
		Code:        http.StatusNotFound,
		Message:     "Not found",
		Description: &description,
		Fields:      nil,
	}
}

func NewErrorByUnauthorized(err error) *handle {
	description := err.Error()
	return &handle{
		Code:        http.StatusUnauthorized,
		Message:     "Unauthorized",
		Description: &description,
		Fields:      nil,
	}
}

func NewErrorByForbidden(err error) *handle {
	description := err.Error()
	return &handle{
		Code:        http.StatusForbidden,
		Message:     "Forbidden",
		Description: &description,
		Fields:      nil,
	}
}

func NewErrorByBadRequest(err error) *handle {
	description := err.Error()
	return &handle{
		Code:        http.StatusBadRequest,
		Message:     "Bad request",
		Description: &description,
		Fields:      nil,
	}
}
