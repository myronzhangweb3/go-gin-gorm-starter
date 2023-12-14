package global

import (
	"net/http"

	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
)

var (
	ServerError = NewError(http.StatusInternalServerError, 500, "系统异常，请稍后重试!")
	NotFound    = NewError(http.StatusNotFound, 404, http.StatusText(http.StatusNotFound))
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
	return NewError(http.StatusBadRequest, 400, message)
}

func NewError(statusCode, Code int, msg string) *Error {
	return &Error{
		StatusCode: statusCode,
		Code:       Code,
		Msg:        msg,
	}
}

func ErrHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var Err *Error
				if e, ok := err.(*Error); ok {
					Err = e
				} else if e, ok := err.(error); ok {
					Err = OtherError(e.Error())
				} else {
					Err = ServerError
				}
				// 记录一个错误的日志
				log.Errorf("http error: %v", Err.Msg)
				c.JSON(Err.StatusCode, Err)
				return
			}
		}()
		c.Next()
	}
}
