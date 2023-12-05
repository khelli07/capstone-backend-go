package events

import (
	"backend-go/models"
	"backend-go/repository"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateEvent(c *gin.Context) {
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

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if body.Name == "" || body.Description == "" || body.Location == "" || body.Price == 0 || body.Capacity == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "name, description, location, price, and capacity must not be empty"})
		return
	}

	if body.Price < 0 || body.Capacity < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "price and capacity must be positive"})
		return
	}

	if body.StartTime == "" || body.EndTime == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing time range"})
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
