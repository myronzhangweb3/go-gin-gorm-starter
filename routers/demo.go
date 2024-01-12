package routers

import (
	"github.com/gin-gonic/gin"
	"go-gin-gorm-starter/controller"
)

func InitStrategyInfoRouter(router *gin.RouterGroup) {
	userInfo := router.Group("/strategy")
	userInfo.POST("/", controller.SaveSimpleStrategy)
	userInfo.GET("/", controller.GetSimpleStrategy)
	userInfo.DELETE("/:id", controller.DeleteSimpleStrategy)
}
