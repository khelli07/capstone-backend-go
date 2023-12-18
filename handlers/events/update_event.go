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
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param id path string true "Event ID"
// @Param body formData payload.UpdateEventRequest true "Event"
// @Success 200 {object} payload.GeneralResponse
// @Router /events/{id} [put]
func UpdateEvent(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var body payload.UpdateEventRequest
	if err := c.ShouldBind(&body); err != nil {
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

	if (body.Lat == 0 && body.Long != 0) || (body.Lat != 0 && body.Long == 0) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Both lat and long must be provided"})
		return
	}

	isOnline := false
	if body.Lat == 0 && body.Long == 0 {
		isOnline = true
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
		StartTime:   startTime,
		EndTime:     endTime,
		Categories:  body.Categories,
		Description: body.Description,
		Price:       body.Price,
		Capacity:    body.Capacity,
		IsOnline:    isOnline,
		Lat:         body.Lat,
		Long:        body.Long,
		Organizer:   body.Organizer,
		DressCode:   body.DressCode,
		AgeLimit:    body.AgeLimit,
	}

	_, err = repository.UpdateEvent(id, &updatedEvent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully",
	})
}
