package categories

import (
	"backend-go/models"
	payload "backend-go/payload/request"
	"backend-go/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateCategory godoc
// @Summary Create a new category
// @Description Create a new category
// @Tags categories
// @Accept  json
// @Produce  json
// @Param body body payload.CreateCategoryRequest true "Create Category"
// @Success 201 {object} payload.CreateResponse
// @Router /categories [post]
func CreateCategory(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var body payload.CreateCategoryRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	category := models.Category{Name: body.Name}
	result, err := repository.CreateCategory(&category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": result.InsertedID})
}
