package repository

import (
	"github.com/abilsabili50/middleware-with-go-fiber/app/model"
	"github.com/abilsabili50/middleware-with-go-fiber/pkg/errs"
)

type IUserRepository interface {
	Create(user model.User) errs.MessageErr
	FindByEmail(email string) (*model.User, errs.MessageErr)
	FindById(userId string) (*model.User, errs.MessageErr)
}
