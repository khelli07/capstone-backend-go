package payload

type CreateReviewRequest struct {
	Rating  int    `json:"rating" binding:"required"`
	Comment string `json:"comment" binding:"required"`
}
