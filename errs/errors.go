// Package errs package provides:
// - struct for handling error's HTTP code and verbose message
// as a single object;
// - methods for the most frequent error’s cases.
package errs

import (
	"net/http"
)

// AppError Represents error's message and HTTP status code.
type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

// AsMessage returns a new error object with omitted HTTP status code.
func (e AppError) AsMessage() *AppError {
	return &AppError{Message: e.Message}
}

// NewNotFoundError Creates an error object with HTTP status code 404, using the given
//string as a message.
func NewNotFoundError(message string) *AppError {
	return &AppError{http.StatusNotFound, message}
}

// NewUnexpectedError Creates an error object with HTTP status code 500, using the given
//string as a message.
func NewUnexpectedError(message string) *AppError {
	return &AppError{http.StatusInternalServerError, message}
}

// NewValidationError Creates an error object with HTTP status code 422, using the given
//string as a message.
func NewValidationError(message string) *AppError {
	return &AppError{http.StatusUnprocessableEntity, message}
}

// NewBadRequestError Creates an error object with HTTP status code 400, using the given
//string as a message.
func NewBadRequestError(message string) *AppError {
	return &AppError{http.StatusBadRequest, message}
}

// NewUnauthorizedError creates an error object with HTTP status code 403, using the given
// string as a message.
func NewUnauthorizedError(message string) *AppError {
	return &AppError{http.StatusForbidden, message}
}
