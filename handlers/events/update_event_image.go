package events

import (
	"backend-go/repository"
	"backend-go/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UploadEventImage godoc
// @Summary Upload an image
// @Description Upload an image
// @Tags events
// @Accept multipart/form-data
// @Produce json
// @Param id path string true "Event ID"
// @Param image formData file true "Image file"
// @Success 200 {object} payload.UploadEventImageResponse
// @Router /events/{id}/image [post]
func UploadEventImage(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	eventID := c.Param("id")
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
	if err := repository.UpdateEventImageURL(eventID, imageUrl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"url": imageUrl})
}
