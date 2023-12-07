package payload

import "backend-go/models"

type GetReviewsResponse struct {
	AverageRating float64         `json:"average_rating"`
	Data          []models.Review `json:"data"`
}
