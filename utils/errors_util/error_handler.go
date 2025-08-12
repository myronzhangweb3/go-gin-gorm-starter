package errors_util

import (
	"fmt"
	"go-gin-gorm-starter/utils/logging"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LogHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		logging.Log.Info("request incoming",
			zap.String("client_ip", c.ClientIP()),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
		)

		c.Next()

		logging.Log.Info("request completed",
			zap.Int("status", c.Writer.Status()),
			zap.Duration("duration", time.Since(start)),
			zap.String("client_ip", c.ClientIP()),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
		)
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
				logging.Log.Error(fmt.Sprintf("http error: %v", Err.Msg))
				c.JSON(Err.StatusCode, Err)
				return
			}
		}()
		c.Next()
	}
}
