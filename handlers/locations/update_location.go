package locations

import (
	"backend-go/models"
	payload "backend-go/payload/request"
	"backend-go/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateLocation godoc
// @Summary Update a location
// @Description Update a location
// @Tags locations
// @Accept  json
// @Produce  json
// @Param id path string true "Location ID"
// @Param body body payload.UpdateLocationRequest true "Update Location"
// @Success 200 {object} models.Location
// @Router /locations/{id} [put]
func UpdateLocation(c *gin.Context) {
	var body payload.UpdateLocationRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	level := repository.ValidateLevel(body.Level)
	if level == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid level"})
		return
	}

	updatedLocation := models.Location{
		Name:  body.Name,
		Level: *level,
	}

	id := c.Param("id")
	_, err := repository.UpdateLocation(id, &updatedLocation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Location updated successfully",
	})
}
