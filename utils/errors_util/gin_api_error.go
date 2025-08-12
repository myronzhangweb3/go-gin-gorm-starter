package errors_util

import (
	"net/http"
)

var (
	ServerError = NewError(http.StatusInternalServerError, 500, "System exception, please try again later!")
)

// Error Structure for error handling
type Error struct {
	StatusCode int    `json:"-"`
	Code       int    `json:"code"`
	Msg        string `json:"msg"`
}

func (e *Error) Error() string {
	return e.Msg
}

func OtherError(message string) *Error {
	return NewError(http.StatusInternalServerError, 500, message)
}

func NewError(statusCode, Code int, msg string) *Error {
	return &Error{
		StatusCode: statusCode,
		Code:       Code,
		Msg:        msg,
	}
}
