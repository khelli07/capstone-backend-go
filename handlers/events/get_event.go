package events

import (
	"backend-go/fs"
	"backend-go/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEventById(c *gin.Context) {
	id := c.Param("id")
	entity, err := repository.GetEventById(fs.CTX, fs.FSClient, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, entity)
}
