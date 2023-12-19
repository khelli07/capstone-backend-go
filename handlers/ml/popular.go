package ml

import (
	"backend-go/models"
	"backend-go/repository"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// GetPopularEvents godoc
// @Summary Get popular events
// @Description Get popular events
// @Tags events
// @Accept  json
// @Produce  json
// @Success 200 {object} payload.GetEventsResponse
// @Router /events/popular [get]
func GetPopularEvents(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	url := os.Getenv("ML_HOST") + "/api/v1/popular"
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error making request to ML service",
		})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error reading response body",
		})
		return
	}

	// Create a map to store the unmarshalled data
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error unmarshalling JSON",
		})
		return
	}

	if resp.StatusCode != 200 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error making request to ML service",
		})
		return
	}

	eventIds := data["data"]
	events := []models.Event{}
	for _, eventId := range eventIds.([]interface{}) {
		event, err := repository.GetEventById(eventId.(string))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		events = append(events, event)
	}

	c.JSON(http.StatusOK, gin.H{"data": events})
}
