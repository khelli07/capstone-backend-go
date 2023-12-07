package events

import (
	"backend-go/models"
	"backend-go/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllEvents godoc
// @Summary Get all events
// @Description Get all events
// @Tags events
// @Accept  json
// @Produce  json
// @Success 200 {object} payload.GetEventsResponse
// @Router /events [get]
func GetAllEvents(c *gin.Context) {
	// TODO: Search filters
	var events []models.Event
	events, err := repository.GetAllEvents()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": events})
}
