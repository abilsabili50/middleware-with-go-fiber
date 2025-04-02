package service

import (
	"log"

	"github.com/abilsabili50/middleware-with-go-fiber/app/dto"
	"github.com/abilsabili50/middleware-with-go-fiber/app/model"
	"github.com/abilsabili50/middleware-with-go-fiber/app/repository"
	"github.com/abilsabili50/middleware-with-go-fiber/pkg/errs"
)

type ITaskService interface {
	CreateTask(payload dto.CreateTaskRequest, userId string) errs.MessageErr
	FindAllPublicTasks() ([]dto.TaskResponse, errs.MessageErr)
	FindAllMyTasks(userId string) ([]dto.TaskResponse, errs.MessageErr)
	FindTaskById(userId, taskId string) (*dto.TaskResponse, errs.MessageErr)
}

type taskService struct {
	taskRepository repository.ITaskRepository
	userService    IUserService
}

func NewTaskService(taskRepository repository.ITaskRepository, userService IUserService) ITaskService {
	return &taskService{
		taskRepository: taskRepository,
		userService:    userService,
	}
}

func (t *taskService) CreateTask(payload dto.CreateTaskRequest, userId string) errs.MessageErr {
	// mapping task from dto into entity
	task := model.NewTask(payload.Title, payload.Desc, userId, payload.IsPublic)

	// perform insert new task
	if err := t.taskRepository.CreateTask(task); err != nil {
		log.Printf("[ERROR] - %v", err)
		return err
	}

	return nil
}
func (t *taskService) FindAllPublicTasks() ([]dto.TaskResponse, errs.MessageErr) {
	// perform get all public tasks
	tasks, err := t.taskRepository.FindAllPublicTask()
	if err != nil {
		log.Printf("[ERROR] - %v", err)
		return nil, err
	}

	// mapping data from entity into dto response
	responses := []dto.TaskResponse{}
	for _, task := range tasks {
		user, _ := t.userService.FindById(task.UserID)
		response := dto.TaskResponse{
			Id:        task.ID,
			Title:     task.Title,
			Desc:      task.Description,
			IsPublic:  task.IsPublic,
			Owner:     user.Name,
			CreatedAt: task.CreatedAt,
			UpdatedAt: task.UpdatedAt,
		}

		responses = append(responses, response)
	}
	return responses, nil
}
func (t *taskService) FindAllMyTasks(userId string) ([]dto.TaskResponse, errs.MessageErr) {
	// perform get all user tasks
	tasks, err := t.taskRepository.FindAllMyTask(userId)
	if err != nil {
		log.Printf("[ERROR] - %v", err)
		return nil, err
	}

	// perform get data user
	user, _ := t.userService.FindById(userId)

	// mapping data from entity into dto response
	responses := []dto.TaskResponse{}
	for _, task := range tasks {
		response := dto.TaskResponse{
			Id:        task.ID,
			Title:     task.Title,
			Desc:      task.Description,
			IsPublic:  task.IsPublic,
			Owner:     user.Name,
			CreatedAt: task.CreatedAt,
			UpdatedAt: task.UpdatedAt,
		}

		responses = append(responses, response)
	}
	return responses, nil
}
func (t *taskService) FindTaskById(userId, taskId string) (*dto.TaskResponse, errs.MessageErr) {
	// perform get data user
	user, _ := t.userService.FindById(userId)

	// perform get data task by task id
	task, err := t.taskRepository.FindTaskById(taskId)
	if err != nil {
		log.Printf("[ERROR] - %v", err)
		return nil, err
	}

	// prevent accessing unpublic task by other users
	if task.UserID != user.ID && !task.IsPublic {
		return nil, errs.NewNotFoundError("task not found")
	}

	// mapping data from entity to dto response
	response := &dto.TaskResponse{
		Id:        task.ID,
		Title:     task.Title,
		Desc:      task.Description,
		IsPublic:  task.IsPublic,
		Owner:     user.Name,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	return response, nil
}
