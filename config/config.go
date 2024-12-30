package config

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct {
		Dsn             string
		MaxIdleConns    int
		MaxOpenConns    int
		ConnMaxLifetime time.Duration
	}
}

var AppConfig *Config

// InitConfig 初始化配置文件
func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
		return
	}
	AppConfig = &Config{}
	err2 := viper.Unmarshal(AppConfig)
	if err2 != nil {
		log.Fatalf("unable to decode into struct, %v", err2)
		return
	}
	InitDB()
}
