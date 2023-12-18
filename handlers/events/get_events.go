package events

import (
	"backend-go/models"
	"backend-go/repository"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// GetAllEvents godoc
// @Summary Get all events
// @Description Get all events
// @Tags events
// @Accept  json
// @Produce  json
// @Param name query string false "Event name"
// @Param categories query string false "Event categories" Example(cat1,cat2)
// @Param price_start query string false "Event price start"
// @Param price_end query string false "Event price end"
// @Param age_limit query string false "Event age limit"
// @Param start_time query string false "Event start time" Example(2023-12-02T20:00:00.000Z)
// @Success 200 {object} payload.GetEventsResponse
// @Router /events [get]
func GetEvents(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	query := bson.M{}
	if c.Query("name") != "" {
		query["name"] = bson.M{"$regex": c.Query("name")}
	}

	if categories := c.Query("categories"); categories != "" {
		query["categories"] = bson.M{"$all": strings.Split(categories, ",")}
	}

	if priceStart := c.Query("price_start"); priceStart != "" {
		f, _ := strconv.ParseFloat(priceStart, 64)
		query["price"] = bson.M{"$gte": f}
	}

	if priceEnd := c.Query("price_end"); priceEnd != "" {
		f, _ := strconv.ParseFloat(priceEnd, 64)
		if existingPrice, exists := query["price"]; exists {
			query["price"] = bson.M{
				"$gte": existingPrice.(bson.M)["$gte"],
				"$lte": f,
			}
		} else {
			query["price"] = bson.M{"$lte": f}
		}
	}

	if ageLimit := c.Query("age_limit"); ageLimit != "" {
		i, _ := strconv.Atoi(ageLimit)
		query["age_limit"] = bson.M{"$lte": i}
	}

	layout := "2006-01-02T15:04:05.000Z"
	if startTime := c.Query("start_time"); startTime != "" {
		t, err := time.Parse(layout, startTime)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		query["start_time"] = bson.M{"$gte": t}
	}

	query["end_time"] = bson.M{"$gte": time.Now()}

	var events []models.Event
	events, err := repository.QueryEvents(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if len(events) == 0 {
		events = []models.Event{}
	}

	c.JSON(http.StatusOK, gin.H{"data": events})
}
