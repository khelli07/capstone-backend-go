package events

import (
	"backend-go/fs"
	"backend-go/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteEvent(c *gin.Context) {
	id := c.Param("id")
	err := repository.DeleteEvent(fs.CTX, fs.FSClient, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
