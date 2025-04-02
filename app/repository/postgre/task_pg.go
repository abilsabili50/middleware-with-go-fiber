package postgre

import (
	"errors"
	"log"

	"github.com/abilsabili50/middleware-with-go-fiber/app/model"
	"github.com/abilsabili50/middleware-with-go-fiber/app/repository"
	"github.com/abilsabili50/middleware-with-go-fiber/pkg/errs"
	"gorm.io/gorm"
)

// declare repository struct
type taskRepository struct {
	db *gorm.DB
}

// declare factory function
func NewTaskRepository(db *gorm.DB) repository.ITaskRepository {
	return &taskRepository{
		db: db,
	}
}

func (t *taskRepository) CreateTask(payload model.Task) errs.MessageErr {
	// perform create new task
	if err := t.db.Create(&payload).Error; err != nil {
		log.Printf("[ERROR] Failed to create task (UserID: %s): %v", payload.UserID, err)
		return errs.NewInternalServerError("failed to create task")
	}
	return nil
}

func (t *taskRepository) FindAllPublicTask() ([]model.Task, errs.MessageErr) {
	// declare tasks variable
	var tasks []model.Task

	// perform retrieving all public tasks
	if err := t.db.Where("is_public = true").Find(&tasks).Error; err != nil {
		log.Println("[ERROR] Failed to fetch public tasks:", err)
		return nil, errs.NewInternalServerError("failed to fetch public tasks")
	}
	return tasks, nil
}

func (t *taskRepository) FindAllMyTask(userId string) ([]model.Task, errs.MessageErr) {
	// declare tasks variable
	var tasks []model.Task

	// perform retrieving all my task
	if err := t.db.Where("user_id = ?", userId).Find(&tasks).Error; err != nil {
		log.Printf("[ERROR] Failed to fetch tasks for UserID: %s - %v", userId, err)
		return nil, errs.NewInternalServerError("failed to fetch your tasks")
	}
	return tasks, nil
}

func (t *taskRepository) FindTaskById(taskId string) (*model.Task, errs.MessageErr) {
	// declare entity variable
	var task model.Task

	// perform retrieving task data by task id
	err := t.db.Where("id = ?", taskId).First(&task).Error
	if err != nil {
		// checker while record < 0
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("[WARN] Task not found (TaskID: %s)", taskId)
			return nil, errs.NewNotFoundError("task not found")
		}
		log.Printf("[ERROR] Database error while fetching task (TaskID: %s) - %v", taskId, err)
		return nil, errs.NewInternalServerError("failed to fetch task")
	}
	return &task, nil
}
