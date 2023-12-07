package locations

import (
	"backend-go/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteLocation godoc
// @Summary Delete a location
// @Description Delete a location
// @Tags locations
// @Accept  json
// @Produce  json
// @Param id path string true "Location ID"
// @Success 200 {object} payload.GeneralResponse
// @Router /locations/{id} [delete]
func DeleteLocation(c *gin.Context) {
	id := c.Param("id")
	err := repository.DeleteLocation(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Location deleted successfully"})
}
