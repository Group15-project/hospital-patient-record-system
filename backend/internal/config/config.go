package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	AppName string
	Port    string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	JWTSecretKey string

	AccessTokenExpiration  time.Duration
	RefreshTokenExpiration time.Duration
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	access, _ := time.ParseDuration(
		viper.GetString("ACCESS_TOKEN_DURATION"),
	)
	refresh, _ := time.ParseDuration(
		viper.GetString("REFRESH_TOKEN_DURATION"),
	)
	return &Config{
		AppName: viper.GetString("APP_NAME"),
		Port:    viper.GetString("PORT"),

		DBHost:     viper.GetString("DB_HOST"),
		DBPort:     viper.GetString("DB_PORT"),
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBName:     viper.GetString("DB_NAME"),

		JWTSecretKey: viper.GetString("JWT_SECRET_KEY"),

		AccessTokenExpiration:  access,
		RefreshTokenExpiration: refresh,
	}, nil
}
