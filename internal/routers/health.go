package routers

import (
	"go-gin-gorm-starter/pkg/gin2"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (r *Router) InitHealthRouter(log *zap.Logger, router *gin.RouterGroup) {
	health := router.Group("/")
	health.GET("/", func(ctx *gin.Context) {
		gin2.HttpResponse(log, ctx, "ok", nil)
	})
}
