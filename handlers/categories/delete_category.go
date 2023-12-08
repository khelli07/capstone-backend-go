package categories

import (
	"backend-go/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteCategory godoc
// @Summary Delete a category
// @Description Delete a category
// @Tags categories
// @Accept  json
// @Produce  json
// @Param id path string true "Category ID"
// @Success 200 {object} payload.GeneralResponse
// @Router /categories/{id} [delete]
func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	err := repository.DeleteCategory(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
