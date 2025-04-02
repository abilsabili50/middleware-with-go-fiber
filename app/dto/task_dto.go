package dto

import "time"

type CreateTaskRequest struct {
	Title    string `json:"title" validate:"required"`
	Desc     string `json:"desc" validate:"required"`
	IsPublic bool   `json:"is_public" validate:"required"`
}

type TaskResponse struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Desc      string    `json:"desc"`
	IsPublic  bool      `json:"is_public"`
	Owner     string    `json:"owner"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
