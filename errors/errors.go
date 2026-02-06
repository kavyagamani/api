package errors

import "net/http"

// APIError is a reusable error type for the application.
// It contains an HTTP status (used by handlers) and a JSON-friendly body.
type APIError struct {
	HTTPStatus int    `json:"-"`
	Code       int    `json:"code"`
	Message    string `json:"message"`
}

// Error implements the error interface.
func (e *APIError) Error() string { return e.Message }

// New creates a new APIError instance.
func New(httpStatus, code int, message string) *APIError {
	return &APIError{HTTPStatus: httpStatus, Code: code, Message: message}
}

// Predefined errors
var (
	ErrMissingLocation = New(http.StatusBadRequest, 1000, "location parameter is required")
)
