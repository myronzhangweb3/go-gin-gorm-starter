package global

import (
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/spf13/viper"
	"go-gin-gorm-starter/models"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"strings"
)

func init() {
	InitLogger()
	InitDB()
}

func InitDB() {
	var (
		err   error
		env   = os.Getenv("GO_ENV")
		level logger.LogLevel
	)
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	dbname := viper.GetString("database.dbname")
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")

	if env != "prod" {
		level = logger.Info
	} else {
		level = logger.Silent
	}

	gormConfig := &gorm.Config{
		QueryFields: true,
		Logger:      logger.Default.LogMode(level),
	}

	log.Infof("driver: %s", viper.GetString("database.driver"))

	switch viper.GetString("database.driver") {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			username, password, host, port, dbname)
		DB, err = gorm.Open(mysql.Open(dsn), gormConfig)
		if err != nil {
			panic(err)
		}
		break
	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			host, username, password, dbname, port)
		DB, err = gorm.Open(postgres.Open(dsn), gormConfig)
		if err != nil {
			panic(err)
		}
	default:
		panic("database driver not support")
	}

	err = DB.AutoMigrate(
		&models.Demo{},
	)
	if err != nil {
		panic("Migrate DB error:" + err.Error())
	}

	if strings.ToUpper(viper.GetString("database.log.level")) == "DEBUG" {
		DB = DB.Debug()
	}

	DB.Set("gorm:table_options", "CHARSET=utf8mb4")
}
