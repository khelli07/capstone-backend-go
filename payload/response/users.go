package payload

import (
	"backend-go/models"
)

type UserResponse struct {
	ID                   string         `json:"id"`
	Username             string         `json:"username"`
	Email                string         `json:"email"`
	Lat                  float64        `json:"lat"`
	Long                 float64        `json:"long"`
	EventCategories      []string       `json:"event_categories"`
	PreferenceCategories []string       `json:"preference_categories"`
	JoinedEvent          []models.Event `json:"joined_event"`
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
