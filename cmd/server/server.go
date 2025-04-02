package server

import (
	"fmt"
	"log"

	"github.com/abilsabili50/middleware-with-go-fiber/app/api/middleware"
	"github.com/abilsabili50/middleware-with-go-fiber/app/api/route"
	"github.com/abilsabili50/middleware-with-go-fiber/cmd/factory"
	"github.com/abilsabili50/middleware-with-go-fiber/pkg/config"
	"github.com/abilsabili50/middleware-with-go-fiber/pkg/validator"
	"github.com/abilsabili50/middleware-with-go-fiber/platform/database"
	"github.com/abilsabili50/middleware-with-go-fiber/platform/migrations"
	"github.com/gofiber/fiber/v2"
)

func Serve() {
	// load config
	config := config.NewConfig()

	// create connection & perform auto migrate
	db := database.NewDatabase(config.DB)
	migrations.AutoMigrate(db.Conn)

	// create fiber app instance
	app := fiber.New(config.Fiber)

	// factory initialization
	validator := validator.NewRequestValidator()
	repo := factory.CreateRepositories(db.Conn)
	service := factory.CreateServices(config, repo)
	controller := factory.CreateControllers(service, validator)
	middleware := middleware.NewMiddleware(config, repo)

	// route initialization
	api := route.NewAPI(app, controller, middleware)
	api.HandleAPI()

	// running server
	address := fmt.Sprintf("%s:%d", config.App.Host, config.App.Port)
	if err := app.Listen(address); err != nil {
		log.Fatalln("error while running applications")
	}
}
