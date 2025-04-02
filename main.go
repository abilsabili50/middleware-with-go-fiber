package main

import (
	"github.com/abilsabili50/middleware-with-go-fiber/cmd/server"
	_ "github.com/abilsabili50/middleware-with-go-fiber/docs"
)

// @title Middleware Implementation API with Go-Fiber
// @version 1.0
// @description Here is a documentation of this API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey BearerAuth
// @in                         header
// @name                       Authorization

// @host localhost:3000
// @BasePath /api/v1
// @schemes http
func main() {
	server.Serve()
}
