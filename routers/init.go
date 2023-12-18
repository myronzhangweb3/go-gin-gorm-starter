package routers

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go-gin-gorm-starter/global"
	"os"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func InitRouter() *gin.Engine {
	router := gin.New()
	if os.Getenv("ENV") == "local" {
		router.Use(CORSMiddleware())
	}
	v1 := router.Group("api/v1")
	v1.Use(global.LogHandler())
	v1.Use(global.ErrHandler())
	pprof.Register(router)
	loadRouter(v1)
	return router
}

func loadRouter(router *gin.RouterGroup) {
	InitHealthRouter(router)
	InitStrategyInfoRouter(router)
}
