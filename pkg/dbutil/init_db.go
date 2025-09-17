package dbutil

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"go-gin-gorm-starter/config"
	"go-gin-gorm-starter/internal/models"
)

func InitDB(config *config.DBConfig) (*gorm.DB, error) {
	username := config.Username
	password := config.Password
	dbname := config.DBName
	host := config.Host
	port := config.Port

	var level logger.LogLevel
	if config.LogLevel != "prod" {
		level = logger.Info
	} else {
		level = logger.Silent
	}

	gormConfig := &gorm.Config{
		QueryFields: true,
		Logger:      logger.Default.LogMode(level),
	}

	fmt.Printf("driver: %s\n", config.Driver)

	var (
		db  *gorm.DB
		err error
	)
	switch config.Driver {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			username, password, host, port, dbname)
		db, err = gorm.Open(mysql.Open(dsn), gormConfig)
		if err != nil {
			return nil, err
		}
	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
			host, username, password, dbname, port)
		db, err = gorm.Open(postgres.Open(dsn), gormConfig)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("database driver not support")
	}

	err = db.AutoMigrate(
		&models.SimpleStrategy{},
	)
	if err != nil {
		return nil, fmt.Errorf("migrate DB error: %v", err.Error())
	}

	db.Debug().Set("gorm:table_options", "CHARSET=utf8mb4")
	return db, nil
}
