package postgre

import (
	"errors"
	"fmt"
	"log"

	"github.com/abilsabili50/middleware-with-go-fiber/app/model"
	"github.com/abilsabili50/middleware-with-go-fiber/app/repository"
	"github.com/abilsabili50/middleware-with-go-fiber/pkg/errs"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.IUserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) Create(user model.User) errs.MessageErr {
	// perform insert new user
	if err := u.db.Create(&user).Error; err != nil {
		log.Printf("[ERROR] Failed to create user (Email: %s) - %v", user.Email, err)
		return errs.NewInternalServerError("failed to create user")
	}
	return nil
}

func (u *userRepository) FindByEmail(email string) (*model.User, errs.MessageErr) {
	// declare entity variable
	var user model.User

	// perform find user by email
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("[WARN] User not found (Email: %s)", email)
			return nil, errs.NewNotFoundError(fmt.Sprintf("user with email %s is not found", email))
		}
		log.Printf("[ERROR] Database error while fetching user (Email: %s) - %v", email, err)
		return nil, errs.NewInternalServerError("failed to fetch user")
	}

	return &user, nil
}

func (u *userRepository) FindById(userId string) (*model.User, errs.MessageErr) {
	// declare entity variable
	var user model.User

	// perform find user by user id
	if err := u.db.Where("id = ?", userId).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("[WARN] User not found (UserID: %s)", userId)
			return nil, errs.NewUnautorizhedError("unauthorized")
		}
		log.Printf("[ERROR] Database error while fetching user (UserID: %s) - %v", userId, err)
		return nil, errs.NewInternalServerError("failed to fetch user")
	}

	return &user, nil
}
