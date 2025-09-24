package gin2

import (
	"net/http"
)

var (
	commonMsg = "The system is busy, please try again later!"

	CommonError = newError(10000, commonMsg)
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

func newError(Code int, msg string) *Error {
	return &Error{
		StatusCode: http.StatusOK,
		Code:       Code,
		Msg:        msg,
	}
}
