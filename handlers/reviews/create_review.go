package reviews

import (
	"backend-go/models"
	payload "backend-go/payload/request"
	"backend-go/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateReview godoc
// @Summary Create a review
// @Description Create a review
// @Tags reviews
// @Accept  x-www-form-urlencoded
// @Param Authorization header string true "With the bearer started"
// @Param event_id path string true "Event ID"
// @Param body formData payload.CreateReviewRequest true "Review"
// @Success 201 {object} payload.CreateResponse
// @Router /reviews/{event_id} [post]
func CreateReview(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var body payload.CreateReviewRequest
	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if body.Rating < 1 || body.Rating > 5 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "rating must be between 1 and 5"})
		return
	}

	userId := c.MustGet("user").(models.TokenUser).ID
	user, err := repository.GetUserById(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	eventId := c.Param("event_id")
	event, err := repository.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	review := models.Review{
		EventID:     eventId,
		UserID:      userId,
		Category:    event.Category,
		Rating:      body.Rating,
		Comment:     body.Comment,
		JoinedEvent: user.JoinedEvent,
	}

	result, err := repository.CreateReview(&review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": result.InsertedID})
}
