package model

import (
	"log"
	"time"

	"github.com/abilsabili50/middleware-with-go-fiber/pkg/util"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"not null;unique"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func NewUser(name, email, password string) (User, error) {
	currTime := time.Now()
	user := User{
		Name:      name,
		Email:     email,
		CreatedAt: currTime,
		UpdatedAt: currTime,
	}

	hashedPassword, err := util.Hash(password)
	if err != nil {
		log.Printf("[ERROR] Server error while hashing password - %v", err)
		return user, err
	}

	user.Password = hashedPassword

	return user, nil
}

func (u *User) Compare(reqPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(reqPassword)); err != nil {
		return false
	}

	return true
}
