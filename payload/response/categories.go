package payload

import "backend-go/models"

type GetCategoriesResponse struct {
	Data []models.Category `json:"data"`
}
