package events

import (
	"backend-go/models"
	"backend-go/repository"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func UpdateEvent(c *gin.Context) {
	var body struct {
		Name        string   `json:"name"`
		Categories  []string `json:"categories"`
		Description string   `json:"description"`
		Location    string   `json:"location"`
		Price       float32  `json:"price"`
		Capacity    int32    `json:"capacity"`
		Organizer   string   `json:"organizer"`
		DressCode   string   `json:"dress_code"`
		AgeLimit    int      `json:"age_limit"`
		StartTime   string   `json:"start_time"`
		EndTime     string   `json:"end_time"`
	}

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
