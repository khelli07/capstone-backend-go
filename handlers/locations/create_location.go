package locations

import (
	"backend-go/models"
	payload "backend-go/payload/request"
	"backend-go/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateLocation godoc
// @Summary Create a new location
// @Description Create a new location
// @Tags locations
// @Accept  json
// @Produce  json
// @Param body body payload.CreateLocationRequest true "Create Location"
// @Success 201 {object} payload.CreateResponse
// @Router /locations [post]
func CreateLocation(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var body payload.CreateLocationRequest
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	level := repository.ValidateLevel(body.Level)
	if level == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid level"})
		return
	}

	location := models.Location{
		Name:  body.Name,
		Level: *level,
	}

	result, err := repository.CreateLocation(&location)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": result.InsertedID})
}
