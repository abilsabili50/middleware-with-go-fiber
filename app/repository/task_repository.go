package repository

import (
	"github.com/abilsabili50/middleware-with-go-fiber/app/model"
	"github.com/abilsabili50/middleware-with-go-fiber/pkg/errs"
)

type ITaskRepository interface {
	CreateTask(payload model.Task) errs.MessageErr
	FindAllPublicTask() ([]model.Task, errs.MessageErr)
	FindAllMyTask(userId string) ([]model.Task, errs.MessageErr)
	FindTaskById(taskId string) (*model.Task, errs.MessageErr)
}
