package gin2

import (
	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	HttpResponse = Response()
)

func Response() func(ctx *gin.Context, result interface{}, err error) {
	return func(ctx *gin.Context, result interface{}, err error) {
		code := http.StatusOK
		message := "success"
		if err != nil {
			code = http.StatusInternalServerError
			message = err.Error()
			log.Errorf("Request error. URL: %s, Details: %s", ctx.Request.URL, err.Error())
		}
		ctx.JSON(code, gin.H{
			"code":    code,
			"result":  result,
			"message": message,
		})
	}
}
