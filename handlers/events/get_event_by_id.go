package events

import (
	"backend-go/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetEventById godoc
// @Summary Get an event by ID
// @Description Get an event by ID
// @Tags events
// @Accept  json
// @Produce  json
// @Param id path string true "Event ID"
// @Success 200 {object} models.Event
// @Router /events/{id} [get]
func GetEventById(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	id := c.Param("id")
	entity, err := repository.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	category, err := repository.GetCategoryById(entity.Category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	entity.Category = category.Name

	c.JSON(http.StatusOK, entity)
}
