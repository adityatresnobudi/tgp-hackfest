package main

import (
	"log"

	"github.com/dinata1312/TechGP-Project/config"
	"github.com/dinata1312/TechGP-Project/internal/server"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("err loading .env file: %s\n", err.Error())
	}
}

func main() {
	cfg := config.NewConfig()

	s := server.NewServer(cfg)

	s.Run()
}
