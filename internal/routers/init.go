package routers

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go-gin-gorm-starter/internal/middleware"
	"go-gin-gorm-starter/utils/errors_util"
	"os"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	if os.Getenv("ENABLE_CORS") == "1" {
		router.Use(middleware.CORSMiddleware())
	}
	v1 := router.Group("api/v1")
	v1.Use(errors_util.LogHandler())
	v1.Use(middleware.ErrHandler())
	pprof.Register(router)
	loadRouter(v1)
	return router
}

func loadRouter(router *gin.RouterGroup) {
	InitHealthRouter(router)
	InitStrategyInfoRouter(router)
}
