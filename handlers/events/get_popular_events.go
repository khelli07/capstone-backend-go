package events

import (
	"backend-go/models"
	"backend-go/repository"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPopularEvents(c *gin.Context) {
	topK := c.Params.ByName("topK")
	if topK == "" {
		topK = "10"
	}
	num, err := strconv.Atoi(topK)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Invalid args topK: %s", topK)})
		return
	}

	var events []models.Event
	events, err = repository.GetPopularEvents(num)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": events})
}
