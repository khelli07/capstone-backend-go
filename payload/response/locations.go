package payload

import "backend-go/models"

type GetLocationsResponse struct {
	Data []models.Location `json:"data"`
}
