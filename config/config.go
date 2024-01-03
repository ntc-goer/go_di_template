package config

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"os"
	"time"
)

type AppEnv string

const (
	DevelopmentLocalAppEnv AppEnv = "development.yaml.local"
	DevelopmentAppEnv      AppEnv = "development.yaml"
	ProductionAppEnv       AppEnv = "production"
)

type MongoDB struct {
	Host     string `json:"host"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type Cors struct {
	AllowOrigins []string `json:"allow_origins"`
	AllowMethods []string `json:"allow_methods"`
}

type Config struct {
	MongoDB MongoDB `json:"mongodb"`
	Cors    Cors    `json:"cors"`
}

func NewCors(c *Config) *gin.HandlerFunc {
	cors := cors.New(cors.Config{
		AllowOrigins:     c.Cors.AllowOrigins,
		AllowMethods:     c.Cors.AllowMethods,
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 15 * time.Second,
	})
	return &cors
}

func loadFileName() (out string) {
	env := os.Getenv("APP_ENV")
	switch env {
	case string(DevelopmentAppEnv):
		out = "development.yaml"
		break
	case string(ProductionAppEnv):
		out = "production"
	default:
		out = "development.local.yaml"
	}
	return
}

func Load() *Config {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	configFileName := loadFileName()
	viper.SetConfigName(configFileName)
	err := viper.ReadInConfig() // Find and read the configs file
	if err != nil {             // Handle errors reading the configs file
		panic(fmt.Errorf("fatal error configs file: %w", err))
	}

	var cfg *Config
	err = viper.Unmarshal(&cfg, func(c *mapstructure.DecoderConfig) {
		c.TagName = "json"
	})
	if err != nil {
		panic(fmt.Errorf("unmashal configs fail: %w", err))
	}
	return cfg
}
