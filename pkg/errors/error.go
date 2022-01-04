package errors

import "fmt"

type AppError struct {
	message string
	reason  Reason
}

func (e *AppError) Error() string {
	return e.message
}

func (e *AppError) ErrorWithReason() string {
	return fmt.Sprintf("%v, Reason: %v", e.message, e.reason)
}

func (e *AppError) Reason() string {
	return string(e.reason)
}

func New(message string, reason Reason) error {
	return &AppError{
		message,
		reason,
	}
}
