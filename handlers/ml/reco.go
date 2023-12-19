package ml

import (
	"backend-go/models"
	"backend-go/repository"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// GetRecommendedEvents godoc
// @Summary Get recommended events
// @Description Get recommended events
// @Tags events
// @Accept  json
// @Produce  json
// @Param Authorization header string true "With the bearer started"
// @Success 200 {object} payload.GetEventsResponse
// @Router /events/reco [get]
func GetRecommendedEvents(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	user := c.MustGet("user").(models.TokenUser)
	payload := []byte(`{
		"user_id": "` + user.ID + `",
		"threshold": 0.1
	}`)

	RECO_URL := os.Getenv("ML_HOST") + "/api/v1/infer/"
	req, err := http.NewRequest("POST", RECO_URL, bytes.NewBuffer(payload))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error making request to ML service"})
		return
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error making request to ML service"})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error reading response from ML service"})
		return
	}

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
