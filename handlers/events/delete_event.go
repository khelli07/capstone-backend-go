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
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Event deleted successfully"})
}
