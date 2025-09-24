package routers

import (
	"go-gin-gorm-starter/internal/controller"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (r *Router) InitStrategyInfoRouter(log *zap.Logger, router *gin.RouterGroup) {
	userInfo := router.Group("/strategy")
	simpleStrategyController := controller.NewSimpleStrategyController(log, r.db)
	userInfo.POST("/", simpleStrategyController.SaveSimpleStrategy)
	userInfo.GET("/", simpleStrategyController.GetSimpleStrategy)
	userInfo.DELETE("/:id", simpleStrategyController.DeleteSimpleStrategy)
}
