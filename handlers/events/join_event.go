package events

import (
	"backend-go/models"
	"backend-go/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JoinEvent godoc
// @Summary Join an event
// @Description Join an event
// @Tags events
// @Accept  json
// @Param Authorization header string true "With the bearer started"
// @Param id path string true "Event ID"
// @Success 200 {object} payload.GeneralResponse
// @Router /events/{id}/join [post]
func JoinEvent(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	eventID := c.Param("id")
	event, err := repository.GetEventById(eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	tokenUser := c.MustGet("user").(models.TokenUser)
	user, err := repository.GetUserById(tokenUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	for _, joinedEvent := range user.JoinedEvent {
		if joinedEvent == eventID {
			c.JSON(http.StatusBadRequest, gin.H{"message": "You already joined this event"})
			return
		}
	}
	user.JoinedEvent = append(user.JoinedEvent, eventID)
	user.EventCategories = append(user.EventCategories, event.Category)

	_, err = repository.UpdateUser(tokenUser.ID, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

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
