package gin2

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func HttpResponse(log *zap.Logger, ctx *gin.Context, result interface{}, err error) {
	bizResponse := &Response{}
	httpCode := 200
	bodyCode := 0
	message := ""
	if err != nil {
		var err2 *Error
		if e, ok := err.(*Error); ok {
			err2 = e
		} else {
			err2 = CommonError
		}
		httpCode = err2.StatusCode
		bodyCode = err2.Code
		message = err2.Msg
		log.Error(fmt.Sprintf("request error. URL: %s, Details: %s", ctx.Request.URL, err.Error()))

		bizResponse.Msg = message
		bizResponse.Code = bodyCode
	} else {
		bizResponse.Data = result
	}
	ctx.JSON(httpCode, bizResponse)
}
