package reviews

import (
	"backend-go/repository"

	"github.com/gin-gonic/gin"
)

func GetReviews(c *gin.Context) {
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
