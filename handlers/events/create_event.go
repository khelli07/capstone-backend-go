package events

import (
	"backend-go/models"
	payload "backend-go/payload/request"
	"backend-go/repository"
	"backend-go/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateEvent godoc
// @Summary Create an event
// @Description Create an event.
// @Tags events
// @Accept  multipart/form-data
// @Produce json
// @Param image formData file false "Image file"
// @Param body formData payload.CreateEventData true "Event"
// @Success 200 {object} payload.CreateResponse
// @Router /events [post]
func CreateEvent(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var body payload.CreateEventRequest
	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Lat and Long
	if (body.Lat == 0 && body.Long != 0) || (body.Lat != 0 && body.Long == 0) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Both lat and long must be provided"})
		return
	}

	isOnline := false
	if body.Lat == 0 && body.Long == 0 {
		isOnline = true
	}

	// Time
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

	// Category
	category, err := repository.GetCategoryByName(body.Category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Image
	f, uploadedFile, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	defer f.Close()

	fileName, err := utils.UploadFile(f, uploadedFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	imageUrl := fmt.Sprintf("%s%s", "https://storage.googleapis.com", fileName)

	var event = models.Event{
		Name:        body.Name,
		StartTime:   startTime,
		EndTime:     endTime,
		Category:    category.ID.Hex(),
		Description: body.Description,
		Price:       body.Price,
		Capacity:    body.Capacity,
		ImageURL:    imageUrl,
		IsOnline:    isOnline,
		Lat:         body.Lat,
		Long:        body.Long,
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
