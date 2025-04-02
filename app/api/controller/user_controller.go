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

// define controller interface
type IUserController interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

// define controller struct
type userController struct {
	validator   validator.RequestValidator
	userService service.IUserService
}

// define factory function
func NewUserController(validator validator.RequestValidator, userService service.IUserService) IUserController {
	return &userController{
		validator:   validator,
		userService: userService,
	}
}

// Register create new user
//
//		@Summary		Register
//		@Description	create new user
//		@Tags				users
//	 	@Accept 		json
//		@Produce		json
//
//		@Param		input body	dto.RegisterRequest	true	"user registration input"
//
//		@Success		201		{object}	dto.Response[any]{}
//		@Failure		400		{object}	dto.Response[any]{data=[]string}
//		@Failure		500		{object}	dto.Response[any]{data=any}
//		@Router			/users/register [post]
func (u *userController) Register(c *fiber.Ctx) error {
	// declare req & res
	var req dto.RegisterRequest
	var res dto.Response[any]

	// parse JSON request to struct
	if err := c.BodyParser(&req); err != nil {
		log.Printf("[ERROR] - %v", err)
		newErrParse := errs.NewBadRequestError("invalid request format")
		res.Mapper("error", newErrParse.Error(), newErrParse.Status(), nil)
		return c.Status(newErrParse.Status()).JSON(res)
	}

	// validate request body
	errsBinding := u.validator.Validate(req)
	if errsBinding != nil {
		log.Printf("[ERROR] - %v", errsBinding)
		newErrBinding := errs.NewBadRequestError("invalid request format")
		res.Mapper("error", newErrBinding.Error(), newErrBinding.Status(), errsBinding)
		return c.Status(newErrBinding.Status()).JSON(res)
	}

	// perform add new user
	if errCreated := u.userService.Register(req); errCreated != nil {
		log.Printf("[ERROR] - %v", errCreated)
		newErrCreated := errs.NewBadRequestError("invalid request format")
		res.Mapper("error", newErrCreated.Error(), newErrCreated.Status(), nil)
		return c.Status(newErrCreated.Status()).JSON(res)
	}

	// return response
	res.Mapper("success", "user created successfully", http.StatusCreated, nil)
	return c.JSON(res)
}

// Login sign in account
//
//		@Summary		Login
//		@Description	login to authenticate
//		@Tags				users
//	 	@Accept 		json
//		@Produce		json
//
//		@Param		input body	dto.LoginRequest	true	"user login input"
//
//		@Success		201		{object}	dto.Response[dto.LoginResponse]{data=dto.LoginResponse}
//		@Failure		400		{object}	dto.Response[any]{data=[]string}
//		@Failure		500		{object}	dto.Response[any]{data=any}
//		@Router			/users/login [post]
func (u *userController) Login(c *fiber.Ctx) error {
	// declare req & res
	var req dto.LoginRequest
	var res dto.Response[any]

	// parse JSON request to struct
	if errParse := c.BodyParser(&req); errParse != nil {
		log.Printf("[ERROR] - %v", errParse)
		newErrParse := errs.NewBadRequestError("invalid request format")
		res.Mapper("error", newErrParse.Error(), newErrParse.Status(), nil)
		return c.Status(newErrParse.Status()).JSON(res)
	}

	// validate request body
	errsBinding := u.validator.Validate(req)
	if errsBinding != nil {
		log.Printf("[ERROR] - %v", errsBinding)
		newErrBinding := errs.NewBadRequestError("invalid request format")
		res.Mapper("error", newErrBinding.Error(), newErrBinding.Status(), errsBinding)
		return c.Status(newErrBinding.Status()).JSON(res)
	}

	// perform login
	response, errLogin := u.userService.Login(req)
	if errLogin != nil {
		log.Printf("[ERROR] - %v", errLogin)
		newErrLogin := errs.NewBadRequestError("invalid request format")
		res.Mapper("error", newErrLogin.Error(), newErrLogin.Status(), nil)
		return c.Status(newErrLogin.Status()).JSON(res)
	}

	// return response
	res.Mapper("success", "login success", http.StatusOK, response)
	return c.JSON(res)
}
