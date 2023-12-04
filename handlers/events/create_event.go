package events

import (
	"backend-go/models"
	"backend-go/repository"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type EventPayload struct {
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

func CreateEvent(c *gin.Context) {
	var eventPayload EventPayload
	if err := c.BindJSON(&eventPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if eventPayload.Name == "" || eventPayload.Description == "" || eventPayload.Location == "" || eventPayload.Price == 0 || eventPayload.Capacity == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "name, description, location, price, and capacity must not be empty"})
		return
	}

	if eventPayload.Price < 0 || eventPayload.Capacity < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "price and capacity must be positive"})
		return
	}

	if eventPayload.StartTime == "" || eventPayload.EndTime == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing time range"})
		return
	}

	layout := "2006-01-02T15:04:05.000Z"
	startTime, err := time.Parse(layout, eventPayload.StartTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid start time"})
		return
	}

	endTime, err := time.Parse(layout, eventPayload.EndTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid end time"})
		return
	}

	if startTime.After(endTime) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "start time must be before end time"})
		return
	}

	var event = models.Event{
		Name:        eventPayload.Name,
		StartTime:   startTime,
		EndTime:     endTime,
		Categories:  eventPayload.Categories,
		Description: eventPayload.Description,
		Location:    eventPayload.Location,
		Price:       eventPayload.Price,
		Capacity:    eventPayload.Capacity,
		Organizer:   eventPayload.Organizer,
		DressCode:   eventPayload.DressCode,
		AgeLimit:    eventPayload.AgeLimit,
		TotalLikes:  0,
		Timestamps: models.Timestamps{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	result, err := repository.CreateEvent(&event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": result.InsertedID})
}
