package categories

import (
	"backend-go/models"
	"backend-go/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCategories godoc
// @Summary Get all categories
// @Description Get all categories
// @Tags categories
// @Accept  json
// @Produce  json
// @Success 200 {object} payload.GetCategoriesResponse
// @Router /categories [get]
func GetCategories(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var categories []models.Category
	categories, err := repository.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if len(categories) == 0 {
		categories = []models.Category{}
	}

	c.JSON(http.StatusOK, gin.H{"data": categories})
}
