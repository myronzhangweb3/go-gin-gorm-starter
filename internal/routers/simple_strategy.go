package routers

import (
	"github.com/gin-gonic/gin"
	"go-gin-gorm-starter/internal/controller"
	"go-gin-gorm-starter/utils/dbutil"
)

func InitStrategyInfoRouter(router *gin.RouterGroup) {
	userInfo := router.Group("/strategy")
	simpleStrategyController := controller.NewSimpleStrategyController(dbutil.DB)
	userInfo.POST("/", simpleStrategyController.SaveSimpleStrategy)
	userInfo.GET("/", simpleStrategyController.GetSimpleStrategy)
	userInfo.DELETE("/:id", simpleStrategyController.DeleteSimpleStrategy)
}
