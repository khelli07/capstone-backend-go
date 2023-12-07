package events

import (
	"backend-go/models"
	payload "backend-go/payload/request"
	"backend-go/repository"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// UpdateEvent godoc
// @Summary Update an event
// @Description Update an event
// @Tags events
// @Accept  json
// @Produce  json
// @Param id path string true "Event ID"
// @Param body body payload.UpdateEventRequest true "Event"
// @Success 200 {object} payload.GeneralResponse
// @Router /events/{id} [put]
func UpdateEvent(c *gin.Context) {
	var body payload.UpdateEventRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	id := c.Param("id")
	_, err := repository.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Event not found",
		})
		return
	}

	layout := "2006-01-02T15:04:05.000Z"

	startTime, err := time.Parse(layout, body.StartTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid start time"})
		return
	}

	endTime, err := time.Parse(layout, body.EndTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid end time"})
		return
	}

	if startTime.After(endTime) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "start time must be before end time"})
		return
	}

	updatedEvent := models.Event{
		Name:        body.Name,
		Categories:  body.Categories,
		Description: body.Description,
		Location:    body.Location,
		Price:       body.Price,
		Capacity:    body.Capacity,
		Organizer:   body.Organizer,
		DressCode:   body.DressCode,
		AgeLimit:    body.AgeLimit,
		StartTime:   startTime,
		EndTime:     endTime,
	}

	_, err = repository.UpdateEvent(id, &updatedEvent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update event",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully",
	})
}
