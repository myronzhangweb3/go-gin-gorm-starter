package config

import "github.com/BurntSushi/toml"

type Config struct {
	Log   *Log         `toml:"log"`
	HTTP  *HTTPConfig  `toml:"http"`
	DB    *DBConfig    `toml:"db"`
	Alarm *AlarmConfig `toml:"alarm"`
}

type HTTPConfig struct {
	Port       int    `toml:"port"`
	EnableCORS bool   `toml:"enable_cors"`
	GinMode    string `toml:"gin_mode"`
}

type Log struct {
	LogLevel string `toml:"log_level"`
}

type DBConfig struct {
	Driver   string `toml:"driver"`
	LogLevel string `toml:"log_level"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	DBName   string `toml:"dbname"`
}

type AlarmConfig struct {
	URL string `toml:"url"`
}

func LoadConfig(path string) (*Config, error) {
	var cfg Config
	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
