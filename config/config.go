package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

var C = make(map[string]interface{}, 0)

func init() {
	// Read in from .env file if available
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Load Config error: %s", err)
	}

	// Read in from environment variables
	_ = viper.BindEnv("PORT")
	_ = viper.BindEnv("DATABASE.DRIVER")
	_ = viper.BindEnv("DATABASE.LOG.LEVEL")
	_ = viper.BindEnv("DATABASE.HOST")
	_ = viper.BindEnv("DATABASE.PORT")
	_ = viper.BindEnv("DATABASE.USERNAME")
	_ = viper.BindEnv("DATABASE.PASSWORD")
	_ = viper.BindEnv("DATABASE.DBNAME")

	viper.Unmarshal(&C)

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		viper.Unmarshal(&C)
		log.Printf("Reload Config success: %s", e.Name)
	})
}
