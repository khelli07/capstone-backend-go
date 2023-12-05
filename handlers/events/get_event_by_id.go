package events

import (
	"backend-go/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEventById(c *gin.Context) {
	id := c.Param("id")
	entity, err := repository.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, entity)
}
