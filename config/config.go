package config

import (
	"os"

	"github.com/dinata1312/TechGP-Project/pkg/constants"
)

type Config struct {
	Postgres PostgresConfig
	Http     HttpConfig
	Jwt      JwtConfig
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
	}

	return cfg
}
