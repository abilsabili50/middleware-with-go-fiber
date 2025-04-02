package config

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	// load godotenv package so we can access .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("can't load .env file. error: %v", err)
	}
}

type Config struct {
	App   *App
	DB    *DBCfg
	Fiber fiber.Config
}

func NewConfig() *Config {
	return &Config{
		App:   LoadApp(),
		DB:    LoadDBCfg(),
		Fiber: LoadFiberConfig(),
	}
}
