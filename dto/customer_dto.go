package dto

import "time"

type CustomerInput struct {
	UserId      int    `json:"user_id" form:"user_id" validate:"required"`
	Name        string `json:"name" form:"name" validate:"required"`
	Email       string `json:"email" form:"email" validate:"required"`
	Gender      string `json:"gender" form:"gender" validate:"required"`
	DateOfBirth string `json:"date_of_birth" form:"date_of_birth" validate:"required"`
	Mobile      string `json:"mobile" form:"mobile" validate:"required"`
	Address     string `json:"address" form:"address" validate:"required"`
}

type CustomerResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CustomerDetailResponse struct {
	Id          int       `json:"id"`
	UserId      int       `json:"user_id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Gender      string    `json:"gender"`
	DateOfBirth string    `json:"date_of_birth"`
	Mobile      string    `json:"mobile"`
	Address     string    `json:"address"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
