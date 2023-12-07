package locations

import (
	"backend-go/models"
	"backend-go/repository"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/exp/slices"

	"github.com/gin-gonic/gin"
)

// GetLocations godoc
// @Summary Get locations
// @Description Get locations
// @Tags locations
// @Accept  json
// @Produce  json
// @Param level query string false "Location level"
// @Success 200 {object} payload.GetLocationsResponse
// @Router /locations [get]
func GetLocations(c *gin.Context) {
	level := c.Query("level")
	query := bson.M{}
	if level != "" {
		levels := []string{"country", "state", "city"}
		if slices.Contains(levels, level) == false {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid level"})
			return
		}

		query["level"] = level
	}

	locations, err := repository.QueryLocation(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if len(locations) == 0 {
		locations = []models.Location{}
	}

	c.JSON(http.StatusOK, gin.H{"data": locations})
	return
}
