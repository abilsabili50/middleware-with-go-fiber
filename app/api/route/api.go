package route

import (
	"github.com/abilsabili50/middleware-with-go-fiber/app/api/middleware"
	"github.com/abilsabili50/middleware-with-go-fiber/cmd/factory"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

// declare API struct
type API struct {
	app        *fiber.App
	controller *factory.Controller
	middleware *middleware.Middleware
}

// declare factory function
func NewAPI(app *fiber.App, controller *factory.Controller, middleware *middleware.Middleware) *API {
	return &API{app: app, controller: controller, middleware: middleware}
}

// declare API http handler
func (a *API) HandleAPI() {
	// handle cors and add logger
	a.app.Use(cors.New())
	a.app.Use(logger.New(logger.Config{Format: "[${time}] [${ip}]:${port} ${status} - ${method} ${path}\n"}))

	// create base path api
	v1 := a.app.Group("/api/v1")

	// create ping path to check that server is running properly
	v1.Get("", func(c *fiber.Ctx) error {
		return c.SendString("server running")
	})

	// swagger api
	v1.Get("/swagger/*", swagger.HandlerDefault)

	// user api
	userRoute := v1.Group("/users")
	userRoute.Post("/register", a.controller.User.Register)
	userRoute.Post("/login", a.controller.User.Login)

	// task api
	taskRoute := v1.Group("/tasks")
	taskRoute.Get("", a.controller.Task.FindAllPublicTask)

	// task api (requires authenticate)
	authorizedRoute := taskRoute.Group("", a.middleware.JWTProtected(), a.middleware.IsRegisteredUser)
	authorizedRoute.Post("", a.controller.Task.CreateTask)
	authorizedRoute.Get("/my", a.controller.Task.FindAllMyTask)
	authorizedRoute.Get("/:taskId", a.controller.Task.FindTaskById)
}
