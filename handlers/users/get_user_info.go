package users

import (
	"backend-go/models"
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
// @Success 200 {object} models.User
// @Router /users [get]
func GetUserInfo(c *gin.Context) {
	user := c.MustGet("user").(models.TokenUser)
	entity, err := repository.GetUserById(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, entity)
}
