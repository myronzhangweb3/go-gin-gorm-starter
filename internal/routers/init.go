package routers

import (
	"fmt"
	"go-gin-gorm-starter/config"
	"go-gin-gorm-starter/internal/middleware"
	"go-gin-gorm-starter/utils/gin2"
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

func (r *Router) InitOrderRouter() *gin.Engine {
	router := gin.New()
	if r.cfg.HTTP.EnableCORS {
		r.log.Info("enable cors")
		router.Use(middleware.CORSMiddleware())
	}
	api := router.Group("api")
	api.Use(r.LogHandler())
	api.Use(r.ErrHandler())
	pprof.Register(router)
	r.InitHealthRouter(api)
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

func (r *Router) ErrHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var Err *gin2.Error
				if e, ok := err.(*gin2.Error); ok {
					Err = e
				} else if e, ok := err.(error); ok {
					Err = gin2.OtherError(e.Error())
				} else {
					Err = gin2.ServerError
				}
				r.log.Error(fmt.Sprintf("http error: %v", Err.Msg))
				c.JSON(Err.StatusCode, Err)
				return
			}
		}()
		c.Next()
	}
}
