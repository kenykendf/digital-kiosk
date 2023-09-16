package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost               string        `mapstructure:"DB_HOST"`
	DBUser               string        `mapstructure:"DB_USER"`
	DBPassword           string        `mapstructure:"DB_PASSWORD"`
	DBName               string        `mapstructure:"DB_NAME"`
	DBPort               string        `mapstructure:"DB_PORT"`
	ServerPort           string        `mapstructure:"SERVER_PORT"`
	LogLevel             string        `mapstructure:"LOG_LEVEL"`
	AccessTokenKey       string        `mapstructure:"ACCESS_TOKEN_KEY"`
	RefreshTokenKey      string        `mapstructure:"REFRESH_TOKEN_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	CloudinaryName       string        `mapstructure:"CLOUDINARY_NAME"`
	CloudinaryAPIKey     string        `mapstructure:"CLOUDINARY_API_KEY"`
	CloudinarySecretKey  string        `mapstructure:"CLOUDINARY_SECRET_KEY"`
	CloudinaryDir        string        `mapstructure:"CLOUDINARY_DIR"`
}

func LoadConfig(fileconfigpath string) (Config, error) {
	var config Config

	viper.AddConfigPath(fileconfigpath)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}
	// nolint:errcheck
	viper.Unmarshal(&config)
	return config, nil
}
