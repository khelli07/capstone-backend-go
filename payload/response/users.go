package payload

import "backend-go/models"

type LoginResponse struct {
	Token string           `json:"token"`
	User  models.TokenUser `json:"user"`
}
