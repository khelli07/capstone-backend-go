package events

import (
	"backend-go/fs"
	"backend-go/models"
	"backend-go/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllEvents(c *gin.Context) {
	// TODO: Search filters
	var events []models.Event
	events, err := repository.GetAllEvents(fs.CTX, fs.FSClient)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": events})
}
