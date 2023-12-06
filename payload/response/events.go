package payload

import "backend-go/models"

type CreateEventResponse struct {
	ID string `json:"id"`
}

type GetEventsResponse struct {
	Data []models.Event `json:"data"`
}
