package controller

import (
	"log"
	"net/http"

	"github.com/abilsabili50/middleware-with-go-fiber/app/dto"
	"github.com/abilsabili50/middleware-with-go-fiber/app/service"
	"github.com/abilsabili50/middleware-with-go-fiber/pkg/errs"
	"github.com/abilsabili50/middleware-with-go-fiber/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

// define interface
type ITaskController interface {
	CreateTask(c *fiber.Ctx) error
	FindAllPublicTask(c *fiber.Ctx) error
	FindAllMyTask(c *fiber.Ctx) error
	FindTaskById(c *fiber.Ctx) error
}

// define struct
type taskController struct {
	validator   validator.RequestValidator
	taskService service.ITaskService
}

// define controller factory
func NewTaskController(validator validator.RequestValidator, taskService service.ITaskService) ITaskController {
	return &taskController{
		validator:   validator,
		taskService: taskService,
	}
}

// CreateTask add new task
//
//		@Summary		CreateTask
//		@Description	add new task (requires authentication)
//		@Tags				tasks
//	 	@Accept 		json
//		@Produce		json
//		@Security		BearerAuth
//
//		@Param		input body	dto.CreateTaskRequest	true	"create task input"
//
//		@Success		201		{object}	dto.Response[any]{data=any}
//		@Failure		400		{object}	dto.Response[any]{data=[]string}
//		@Failure		401		{object}	dto.Response[any]{data=any}
//		@Failure		500		{object}	dto.Response[any]{data=any}
//		@Router			/tasks [post]
func (t *taskController) CreateTask(c *fiber.Ctx) error {
	// declare req & res
	var req dto.CreateTaskRequest
	var res dto.Response[any]

	// catch userId from middleware
	userId := c.Locals("userId").(string)

	// parse JSON request to struct
	if err := c.BodyParser(&req); err != nil {
		log.Printf("[ERROR] - %v", err)
		newErrParse := errs.NewBadRequestError("invalid request format")
		res.Mapper("error", newErrParse.Error(), newErrParse.Status(), nil)
		return c.Status(newErrParse.Status()).JSON(res)
	}

	// validate request body
	errsBinding := t.validator.Validate(req)
	if errsBinding != nil {
		log.Printf("[ERROR] - %v", errsBinding)
		newErrBinding := errs.NewBadRequestError("invalid request format")
		res.Mapper("error", newErrBinding.Error(), newErrBinding.Status(), errsBinding)
		return c.Status(newErrBinding.Status()).JSON(res)
	}

	// perform add new user
	if errCreated := t.taskService.CreateTask(req, userId); errCreated != nil {
		log.Printf("[ERROR] - %v", errCreated)
		res.Mapper("error", errCreated.Error(), errCreated.Status(), nil)
		return c.Status(errCreated.Status()).JSON(res)
	}

	// return response
	res.Mapper("success", "user created successfully", http.StatusCreated, nil)
	return c.JSON(res)
}

// FindPublicTasks find all public tasks
//
//	@Summary		FindPublicTasks
//	@Description	find all public tasks
//	@Tags				tasks
//	@Produce		json
//
//	@Success		200		{object}	dto.Response[[]dto.TaskResponse]{data=[]dto.TaskResponse}
//	@Failure		500		{object}	dto.Response[any]{data=any}
//	@Router			/tasks [get]
func (t *taskController) FindAllPublicTask(c *fiber.Ctx) error {
	// declare response instance
	var res dto.Response[any]

	// perform retrieve all tasks
	tasks, errs := t.taskService.FindAllPublicTasks()
	if errs != nil {
		log.Printf("[ERROR] - %v", errs)
		res.Mapper("error", errs.Error(), errs.Status(), nil)
		return c.Status(errs.Status()).JSON(res)
	}

	// mapping response
	res.Mapper("success", "all public tasks retrieved successfully", http.StatusOK, tasks)
	return c.JSON(res)
}

// FindMyTasks find all my tasks
//
//	@Summary		FindMyTasks
//	@Description	find all my tasks (requires authentication)
//	@Tags				tasks
//	@Produce		json
//	@Security		BearerAuth
//
//	@Success		200		{object}	dto.Response[[]dto.TaskResponse]{data=[]dto.TaskResponse}
//	@Failure		401		{object}	dto.Response[any]{data=any}
//	@Router			/tasks/my [get]
func (t *taskController) FindAllMyTask(c *fiber.Ctx) error {
	// declare response instance
	var res dto.Response[any]

	// retrieve userId
	userId := c.Locals("userId").(string)

	// perform retrieving all tasks
	tasks, errs := t.taskService.FindAllMyTasks(userId)
	if errs != nil {
		log.Printf("[ERROR] - %v", errs)
		res.Mapper("error", errs.Error(), errs.Status(), nil)
		return c.Status(errs.Status()).JSON(res)
	}

	// mapping response
	res.Mapper("success", "all your tasks retrieved successfully", http.StatusOK, tasks)
	return c.JSON(res)
}

// FindTaskById find task by task id
//
//	@Summary		FindTaskById
//	@Description	find task by task id (requires authentication)
//	@Tags				tasks
//	@Produce		json
//	@Security		BearerAuth
//
//	@Param		taskId	path	string	true	"task id"
//
//	@Success		201		{object}	dto.Response[dto.TaskResponse]{data=dto.TaskResponse}
//	@Failure		400		{object}	dto.Response[any]{data=[]string}
//	@Failure		401		{object}	dto.Response[any]{data=any}
//	@Failure		500		{object}	dto.Response[any]{data=any}
//	@Router			/tasks/{taskId} [get]
func (t *taskController) FindTaskById(c *fiber.Ctx) error {
	// declare response instance
	var res dto.Response[any]

	// retrieve userId and taskId
	userId := c.Locals("userId").(string)
	taskId := c.Params("taskId")

	// perform retrieving all tasks
	tasks, errs := t.taskService.FindTaskById(userId, taskId)
	if errs != nil {
		log.Printf("[ERROR] - %v", errs)
		res.Mapper("error", errs.Error(), errs.Status(), nil)
		return c.Status(errs.Status()).JSON(res)
	}

	// mapping response
	res.Mapper("success", "all public tasks retrieved successfully", http.StatusOK, tasks)
	return c.JSON(res)
}
