package run

import (
	"context"
	"errors"
	"fmt"
	"go-gin-gorm-starter/config"
	"go-gin-gorm-starter/utils/dbutil"
	"go-gin-gorm-starter/utils/logging"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouterInitFunc func(log *zap.Logger, cfg *config.Config, db *gorm.DB) *gin.Engine

func HttpServer(configPath, serviceName string, initRouter RouterInitFunc) error {
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		return err
	}

	log, err := logging.InitLogger(serviceName, cfg.Log.LogLevel)
	if err != nil {
		return err
	}
	log.Debug(fmt.Sprintf("config: %+v", cfg))

	db, err := dbutil.InitDB(cfg.DB)
	if err != nil {
		return err
	}

	if cfg.HTTP.GinMode != "prod" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := initRouter(log, cfg, db)

	addr := fmt.Sprintf(":%d", cfg.HTTP.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Server forced to shutdown: ", err)
	}
	fmt.Println("Server exiting")

	return nil
}
