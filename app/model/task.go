package model

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID          string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	IsPublic    bool   `gorm:"not null"`
	UserID      string `gorm:"type:uuid"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func NewTask(title, desc, userId string, isPublic bool) Task {
	currTime := time.Now()
	return Task{
		Title:       title,
		Description: desc,
		IsPublic:    isPublic,
		UserID:      userId,
		CreatedAt:   currTime,
		UpdatedAt:   currTime,
	}
}
