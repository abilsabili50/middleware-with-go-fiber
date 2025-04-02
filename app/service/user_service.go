package service

import (
	"log"

	"github.com/abilsabili50/middleware-with-go-fiber/app/dto"
	"github.com/abilsabili50/middleware-with-go-fiber/app/model"
	"github.com/abilsabili50/middleware-with-go-fiber/app/repository"
	"github.com/abilsabili50/middleware-with-go-fiber/pkg/config"
	"github.com/abilsabili50/middleware-with-go-fiber/pkg/errs"
	"github.com/abilsabili50/middleware-with-go-fiber/pkg/util"
)

type IUserService interface {
	Register(payload dto.RegisterRequest) errs.MessageErr
	Login(payload dto.LoginRequest) (*dto.LoginResponse, errs.MessageErr)
	FindById(userId string) (*model.User, errs.MessageErr)
}

type userService struct {
	userRepository repository.IUserRepository
	config         *config.Config
}

func NewUserService(userRepository repository.IUserRepository, config *config.Config) IUserService {
	return &userService{
		userRepository: userRepository,
		config:         config,
	}
}

func (u *userService) Register(payload dto.RegisterRequest) errs.MessageErr {
	// mapping data from dto request to entity
	user, err := model.NewUser(payload.Name, payload.Email, payload.Password)
	if err != nil {
		log.Printf("[ERROR] Server error while create new user instances - %v", err)
		return errs.NewInternalServerError("failed to create new user instance")
	}

	// perform insert new user
	return u.userRepository.Create(user)
}

func (u *userService) Login(payload dto.LoginRequest) (*dto.LoginResponse, errs.MessageErr) {
	// perform get data user by email
	user, err := u.userRepository.FindByEmail(payload.Email)
	if err != nil {
		log.Printf("[ERROR] - %v", err.Error())
		return nil, err
	}

	// compare db password with req password
	if !user.Compare(payload.Password) {
		log.Printf("[ERROR] - invalid password")
		return nil, errs.NewBadRequestError("invalid credentials")
	}

	// generate token using user id as claims data
	token, err := util.GenerateToken(u.config.App, user.ID)
	if err != nil {
		return nil, err
	}

	// mapping data response
	response := &dto.LoginResponse{
		Type:  "Bearer",
		Token: token,
	}

	return response, nil
}

func (u *userService) FindById(userId string) (*model.User, errs.MessageErr) {
	// perform get data user by user id
	user, err := u.userRepository.FindById(userId)
	if err != nil {
		log.Printf("[ERROR] - %v", err)
		return nil, err
	}

	return user, nil
}
