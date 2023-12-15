package payload

import "backend-go/models"

type GetEventsResponse struct {
	Data []models.Event `json:"data"`
}

type UploadImageResponse struct {
	URL string `json:"url"`
}
