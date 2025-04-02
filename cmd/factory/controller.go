package factory

import (
	"github.com/abilsabili50/middleware-with-go-fiber/app/api/controller"
	"github.com/abilsabili50/middleware-with-go-fiber/pkg/validator"
)

type Controller struct {
	User controller.IUserController
	Task controller.ITaskController
}

// create factory function
func CreateControllers(services *Service, validator validator.RequestValidator) *Controller {
	return &Controller{
		User: controller.NewUserController(validator, services.User),
		Task: controller.NewTaskController(validator, services.Task),
	}
}
