package middleware

import (
	"fmt"
	"go-gin-gorm-starter/utils/errors_util"
	"go-gin-gorm-starter/utils/logging"

	"github.com/gin-gonic/gin"
)

func ErrHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var Err *errors_util.Error
				if e, ok := err.(*errors_util.Error); ok {
					Err = e
				} else if e, ok := err.(error); ok {
					Err = errors_util.OtherError(e.Error())
				} else {
					Err = errors_util.ServerError
				}
				logging.Log.Error(fmt.Sprintf("http error: %v", Err.Msg))
				c.JSON(Err.StatusCode, Err)
				return
			}
		}()
		c.Next()
	}
}
