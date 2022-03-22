package model

import "time"

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUserResponse struct {
	Id        uint      `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
