package utils

import "net/http"

type AppError struct {
	Code    int
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Message
}

func NewAppError(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

var ErrNotFound = NewAppError(http.StatusNotFound, "resource not found", nil)
var ErrBadRequest = NewAppError(http.StatusBadRequest, "bad request", nil)