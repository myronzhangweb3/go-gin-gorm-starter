package routers

import (
	"go-gin-gorm-starter/config"
	"go-gin-gorm-starter/internal/middleware"
	"time"

	"go.uber.org/zap"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Router struct {
	log *zap.Logger
	cfg *config.Config
	db  *gorm.DB
}

func NewRouter(log *zap.Logger, cfg *config.Config, db *gorm.DB) *Router {
	return &Router{log: log, cfg: cfg, db: db}
}

func (r *Router) InitRouter() *gin.Engine {
	router := gin.New()
	if r.cfg.HTTP.EnableCORS {
		r.log.Info("enable cors")
		router.Use(middleware.CORSMiddleware())
	}
	api := router.Group("api")
	api.Use(r.LogHandler())
	pprof.Register(router)
	r.InitHealthRouter(r.log, api)
	r.InitStrategyInfoRouter(r.log, api)
	return router
}

func (r *Router) LogHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		r.log.Debug("request incoming",
			zap.String("client_ip", c.ClientIP()),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
		)

		c.Next()

		r.log.Debug("request completed",
			zap.Int("status", c.Writer.Status()),
			zap.Duration("duration", time.Since(start)),
			zap.String("client_ip", c.ClientIP()),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
		)
	}
}
