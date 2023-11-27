package events

import (
	"backend-go/ds"
	"backend-go/models"
	"backend-go/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateEvent(c *gin.Context) {
	var event models.Event
	if err := c.BindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	key, err := repository.CreateEvent(ds.CTX, ds.DS, &event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": key.ID, "name": key.Name})
}
