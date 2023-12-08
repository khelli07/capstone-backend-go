package events

import (
	"backend-go/models"
	payload "backend-go/payload/request"
	"backend-go/repository"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateEvent godoc
// @Summary Create an event
// @Description Create an event
// @Tags events
// @Accept json
// @Produce json
// @Param body body payload.CreateEventRequest true "Event object"
// @Success 200 {object} payload.CreateResponse
// @Router /events [post]
func CreateEvent(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var body payload.CreateEventRequest
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
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

	var event = models.Event{
		Name:        body.Name,
		StartTime:   startTime,
		EndTime:     endTime,
		Categories:  body.Categories,
		Description: body.Description,
		Location:    body.Location,
		Price:       body.Price,
		Capacity:    body.Capacity,
		Organizer:   body.Organizer,
		DressCode:   body.DressCode,
		AgeLimit:    body.AgeLimit,
		TotalLikes:  0,
	}

	result, err := repository.CreateEvent(&event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": result.InsertedID})
}
