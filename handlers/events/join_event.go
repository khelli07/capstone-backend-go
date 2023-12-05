package events

import (
	"backend-go/models"
	"backend-go/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
