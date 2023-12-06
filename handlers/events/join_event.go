package events

import (
	"backend-go/models"
	"backend-go/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JoinEvent godoc
// @Summary Join an event [LOGIN REQUIRED]
// @Description Join an event
// @Tags events
// @Accept  json
// @Param id path string true "Event ID"
// @Success 200 {object} payload.GeneralResponse
// @Router /events/{id}/join [post]
func JoinEvent(c *gin.Context) {
	eventID := c.Param("id")
	event, err := repository.GetEventById(eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	tokenUser := c.MustGet("user").(models.TokenUser)
	event.Participants = append(event.Participants, tokenUser.ID)
	_, err = repository.UpdateEvent(eventID, &event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully joined event",
	})
}
