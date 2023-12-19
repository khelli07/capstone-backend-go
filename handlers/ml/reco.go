package ml

import (
	"github.com/gin-gonic/gin"
)

func GetRecommendedEvents(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "API not implemented yet",
	})
}
