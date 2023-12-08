package users

import (
	"backend-go/models"
	payload "backend-go/payload/request"
	"backend-go/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateUser godoc
// @Summary Update a user
// @Description Update a user
// @Tags users
// @Accept  json
// @Produce  json
// @Param Authorization header string true "With the bearer started"
// @Param body body payload.UpdateUserRequest true "User"
// @Success 200 {object} payload.GeneralResponse
// @Router /users [put]
func UpdateUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var body payload.UpdateUserRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	tokenUser := c.MustGet("user").(models.TokenUser)
	user, err := repository.GetUserById(tokenUser.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}

	updatedUser := models.User{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
		// Changed
		Username: body.Username,
	}

	_, err = repository.UpdateUser(tokenUser.ID, &updatedUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
	})
}
