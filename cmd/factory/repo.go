package factory

import (
	"github.com/abilsabili50/middleware-with-go-fiber/app/repository"
	"github.com/abilsabili50/middleware-with-go-fiber/app/repository/postgre"
	"gorm.io/gorm"
)

type Repository struct {
	User repository.IUserRepository
	Task repository.ITaskRepository
}

// create factory function
func CreateRepositories(db *gorm.DB) *Repository {
	return &Repository{
		User: postgre.NewUserRepository(db),
		Task: postgre.NewTaskRepository(db),
	}
}
