package routers

import (
	"github.com/gin-gonic/gin"
	"go-gin-gorm-starter/pkg/gin2"
)

func (r *Router) InitHealthRouter(router *gin.RouterGroup) {
	health := router.Group("/")
	health.GET("/", func(ctx *gin.Context) {
		gin2.HttpResponse(ctx, "ok", nil)
	})
}
