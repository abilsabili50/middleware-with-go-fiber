package config

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	Host string
	Port int

	// JWT
	JWTSigningMethod     string
	JWTAccTokenSecretKey string
	JWTAccTokenDuration  time.Duration
}

func LoadApp() *App {
	var app = &App{}

	// read app config
	app.Host = os.Getenv("APP_HOST")
	app.Port, _ = strconv.Atoi(os.Getenv("APP_PORT"))

	app.JWTSigningMethod = os.Getenv("JWT_SIGNING_METHOD")
	app.JWTAccTokenSecretKey = os.Getenv("JWT_ACC_SECRET_KEY")
	accDuration, _ := strconv.Atoi(os.Getenv("JWT_ACC_DURATION_MINUTES"))
	app.JWTAccTokenDuration = time.Duration(accDuration) * time.Minute

	return app
}

func LoadFiberConfig() fiber.Config {
	// load fiber config
	readTimeOut, _ := strconv.Atoi(os.Getenv("APP_READ_TIMEOUT"))
	writeTimeout, _ := strconv.Atoi(os.Getenv("APP_WRITE_TIMEOUT"))
	idleTimeout, _ := strconv.Atoi(os.Getenv("APP_IDLE_TIMEOUT"))
	return fiber.Config{
		ReadTimeout:  time.Second * time.Duration(readTimeOut),
		WriteTimeout: time.Second * time.Duration(writeTimeout),
		IdleTimeout:  time.Second * time.Duration(idleTimeout),
	}
}
