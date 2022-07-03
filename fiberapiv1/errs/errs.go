package errs

import "net/http"

type AppError struct {
	Code    int
	Message string
}

func (e AppError) Error() string {
	return e.Message
}

func NewError(message string) error {
	return AppError{
		Message: message,
	}
}
func NewNotFoundError(message string) error {
	return AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewUnexpectedError() error {
	return AppError{
		Code:    http.StatusInternalServerError,
		Message: "unexpected error",
	}
}
func NewValidationError(message string) error {
	return AppError{
		Code:    http.StatusUnprocessableEntity,
		Message: message,
	}
}

//auth
func NewUnauthorizedError() error {
	return AppError{
		Code:    http.StatusUnauthorized,
		Message: "unauthorized",
	}
}
func NewInvalidAuthTokenError() error {
	return AppError{
		Code:    http.StatusUnauthorized,
		Message: "invalid auth-token",
	}
}
