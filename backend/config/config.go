package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"postgres"`
	JWT      JWTConfig      `mapstructure:"jwt"`
}

type ServerConfig struct {
	Port int
}

type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type JWTConfig struct {
	Secret string
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigName("config")    // 配置文件名稱 (不包含副檔名)
	viper.SetConfigType("json")      // 配置文件類型
	viper.AddConfigPath("../config") // 配置文件所在路徑

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
}
