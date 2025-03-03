package config

import (
	"os"

	"github.com/dinata1312/TechGP-Project/pkg/constants"
)

type Config struct {
	Postgres PostgresConfig
	Http     HttpConfig
	Jwt      JwtConfig
	App      AppConfig
}

type PostgresConfig struct {
	Port     string
	Host     string
	User     string
	Password string
	DBName   string
}

type HttpConfig struct {
	Port string
	Host string
}

type JwtConfig struct {
	SecretKey string
}

type AppConfig struct {
	OriginDomain string
}

func NewConfig() Config {
	cfg := Config{
		Http: HttpConfig{
			Port: os.Getenv(constants.HTTPPort),
			Host: os.Getenv(constants.APIHost),
		},
		Postgres: PostgresConfig{
			Port:     os.Getenv(constants.DBPort),
			Host:     os.Getenv(constants.DBHost),
			User:     os.Getenv(constants.DBUser),
			Password: os.Getenv(constants.DBPassword),
			DBName:   os.Getenv(constants.DBName),
		},
		Jwt: JwtConfig{
			SecretKey: os.Getenv(constants.SecretKey),
		},
		App: AppConfig{
			OriginDomain: os.Getenv(constants.APIHost),
		},
	}

	return cfg
}
