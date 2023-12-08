package events

import (
	"backend-go/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteEvent godoc
// @Summary Delete an event
// @Description Delete an event
// @Tags events
// @Accept  json
// @Produce  json
// @Param id path string true "Event ID"
// @Success 200 {object} payload.GeneralResponse
// @Router /events/{id} [delete]
func DeleteEvent(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	id := c.Param("id")
	err := repository.DeleteEvent(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
