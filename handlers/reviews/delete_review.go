package reviews

import (
	"backend-go/models"
	"backend-go/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteReview godoc
// @Summary Delete a review
// @Description Delete a review
// @Tags reviews
// @Accept  json
// @Param Authorization header string true "With the bearer started"
// @Param id path string true "Review ID"
// @Success 200 {object} payload.GeneralResponse
// @Router /reviews/{id} [delete]
func DeleteReview(c *gin.Context) {
	id := c.Param("id")
	review, err := repository.GetReviewById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if review.UserID != c.MustGet("user").(models.TokenUser).ID {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to delete this review"})
		return
	}

	err = repository.DeleteReview(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Review deleted successfully"})
}
