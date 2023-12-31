package users

import (
	"backend-go/models"
	payload "backend-go/payload/response"
	"backend-go/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserInfo godoc
// @Summary Get user info
// @Description Get user info
// @Tags users
// @Accept  json
// @Produce  json
// @Param Authorization header string true "With the bearer started"
// @Success 200 {object} payload.GetUserResponse
// @Router /users [get]
func GetUserInfo(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	user := c.MustGet("user").(models.TokenUser)
	entity, err := repository.GetUserById(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	categories := []string{}
	for _, categoryID := range entity.PreferenceCategories {
		category, err := repository.GetCategoryById(categoryID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		categories = append(categories, category.Name)
	}

	eventCategories := []string{}
	for _, categoryID := range entity.EventCategories {
		category, err := repository.GetCategoryById(categoryID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		eventCategories = append(eventCategories, category.Name)
	}

	events := []models.Event{}
	for _, eventID := range entity.JoinedEvent {
		event, err := repository.GetEventById(eventID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		events = append(events, event)
	}

	userResponse := payload.UserResponse{
		ID:                   entity.ID.Hex(),
		Username:             entity.Username,
		Email:                entity.Email,
		Lat:                  entity.Lat,
		Long:                 entity.Long,
		EventCategories:      eventCategories,
		PreferenceCategories: categories,
		JoinedEvent:          events,
		Timestamps:           entity.Timestamps,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User fetched successfully",
		"user":    userResponse,
	})
}
