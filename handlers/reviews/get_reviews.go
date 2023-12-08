package reviews

import (
	"backend-go/repository"

	"github.com/gin-gonic/gin"
)

// GetReviews godoc
// @Summary Get reviews of an event
// @Description Get reviews of an event
// @Tags reviews
// @Accept  json
// @Param id path string true "Event ID"
// @Produce  json
// @Success 200 {object} payload.GetReviewsResponse
// @Router /reviews/{event_id} [get]
func GetReviews(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	id := c.Param("event_id")

	reviews, err := repository.GetReviews(id)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	if len(reviews) == 0 {
		c.JSON(200, gin.H{"average_rating": 0, "data": reviews})
		return
	}

	sum := 0
	for _, review := range reviews {
		sum += review.Rating
	}
	averageRating := float64(sum) / float64(len(reviews))

	c.JSON(200, gin.H{"average_rating": averageRating, "data": reviews})
}
