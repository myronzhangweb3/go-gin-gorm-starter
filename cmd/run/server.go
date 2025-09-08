package run

import (
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"go-gin-gorm-starter/config"
	"go-gin-gorm-starter/internal/routers"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Server(ctx *cli.Context) error {
	configPath := ctx.String("config")
	return HttpServer(configPath, "Server", func(log *zap.Logger, cfg *config.Config, db *gorm.DB) *gin.Engine {
		router := routers.NewRouter(log, cfg, db)
		return router.InitOrderRouter()
	})
}
