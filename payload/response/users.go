package payload

import (
	"backend-go/models"
)

type UserResponse struct {
	ID          string         `json:"id"`
	Username    string         `json:"username"`
	Email       string         `json:"email"`
	JoinedEvent []models.Event `json:"joined_event"`
	models.Timestamps
}

type LoginResponse struct {
	Token string           `json:"token"`
	User  models.TokenUser `json:"user"`
}

type GetUserResponse struct {
	Message string       `json:"message"`
	User    UserResponse `json:"user"`
}
