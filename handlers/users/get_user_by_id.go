package users

import (
	"backend-go/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserById godoc
// @Summary Get user by id
// @Description Get user by id
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
func GetUserById(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	id := c.Param("id")
	entity, err := repository.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, entity)
}
