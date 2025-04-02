package factory

import (
	"github.com/abilsabili50/middleware-with-go-fiber/app/service"
	"github.com/abilsabili50/middleware-with-go-fiber/pkg/config"
)

type Service struct {
	User service.IUserService
	Task service.ITaskService
}

// create factory function
func CreateServices(config *config.Config, repository *Repository) *Service {
	userService := service.NewUserService(repository.User, config)
	taskService := service.NewTaskService(repository.Task, userService)
	return &Service{
		User: userService,
		Task: taskService,
	}
}
