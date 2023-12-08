package categories

import (
	"backend-go/models"
	payload "backend-go/payload/request"
	"backend-go/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateCategory godoc
// @Summary Update a category
// @Description Update a category
// @Tags categories
// @Accept  json
// @Produce  json
// @Param id path string true "Category ID"
// @Param body body payload.UpdateCategoryRequest true "Update Category"
// @Success 200 {object} payload.GeneralResponse
// @Router /categories/{id} [put]
func UpdateCategory(c *gin.Context) {
	var body payload.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	updatedCategory := models.Category{Name: body.Name}

	id := c.Param("id")
	_, err := repository.UpdateCategory(id, &updatedCategory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Category updated successfully",
	})
}
