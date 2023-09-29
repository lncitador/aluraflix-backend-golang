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

type errorHandle struct {
	Code        int       `json:"Code"`
	Message     string    `json:"Message"`
	Description *string   `json:"Description,omitempty"`
	Fields      *[]string `json:"fields,omitempty"`
}

func (e errorHandle) Error() string {
	return e.Message
}

func (e errorHandle) Status() int {
	return e.Code
}

func (e errorHandle) GetDescription() string {
	if e.Description != nil {
		return *e.Description
	}

	return ""
}

func NewError(code int, message string, description string) *errorHandle {
	return &errorHandle{
		Code:        code,
		Message:     message,
		Description: &description,
		Fields:      nil,
	}
}

func NewErrorByInternal(err error) *errorHandle {
	description := err.Error()
	return &errorHandle{
		Code:        http.StatusInternalServerError,
		Message:     "Internal error",
		Description: &description,
		Fields:      nil,
	}
}

func NewErrorByValidation(err error) *errorHandle {
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
		return &errorHandle{
			Code:        http.StatusUnprocessableEntity,
			Message:     "Validation error",
			Description: &description,
			Fields:      nil,
		}
	}

	return &errorHandle{
		Code:        http.StatusUnprocessableEntity,
		Message:     "Validation error",
		Description: nil,
		Fields:      &errs,
	}
}

func NewErrorByNotFound(err error) *errorHandle {
	description := err.Error()
	return &errorHandle{
		Code:        http.StatusNotFound,
		Message:     "Not found",
		Description: &description,
		Fields:      nil,
	}
}

func NewErrorByUnauthorized(err error) *errorHandle {
	description := err.Error()
	return &errorHandle{
		Code:        http.StatusUnauthorized,
		Message:     "Unauthorized",
		Description: &description,
		Fields:      nil,
	}
}

func NewErrorByForbidden(err error) *errorHandle {
	description := err.Error()
	return &errorHandle{
		Code:        http.StatusForbidden,
		Message:     "Forbidden",
		Description: &description,
		Fields:      nil,
	}
}

func NewErrorByBadRequest(err error) *errorHandle {
	description := err.Error()
	return &errorHandle{
		Code:        http.StatusBadRequest,
		Message:     "Bad request",
		Description: &description,
		Fields:      nil,
	}
}
