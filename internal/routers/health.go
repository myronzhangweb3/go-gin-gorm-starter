package routers

import (
	"github.com/gin-gonic/gin"
)

func InitHealthRouter(router *gin.RouterGroup) {
	health := router.Group("/")
	health.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ok",
		})
	})
}
