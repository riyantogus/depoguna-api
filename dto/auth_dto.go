package dto

import "time"

type RegisterInput struct {
	Name                 string `json:"name" form:"name" validate:"required"`
	Email                string `json:"email" form:"email" validate:"required,email"`
	Password             string `json:"password" form:"password" validate:"required,min=8"`
	PasswordConfirmation string `json:"password_confirmation" form:"password_confirmation" validate:"required,eqfield=Password"`
}

type RegisterResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginInput struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}

type LoginResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}
