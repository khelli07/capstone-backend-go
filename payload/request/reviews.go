package payload

type CreateReviewRequest struct {
	Rating  int    `form:"rating" binding:"required"`
	Comment string `form:"comment" binding:"required"`
}
