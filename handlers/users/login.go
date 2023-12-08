package users

import (
	payload "backend-go/payload/request"
	"backend-go/repository"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Login godoc
// @Summary Login
// @Description Login
// @Tags users
// @Accept  json
// @Produce  json
// @Param body body payload.LoginRequest true "Login"
// @Success 200 {object} payload.LoginResponse
// @Router /auth/login [post]
func Login(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var body payload.LoginRequest
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to process request body",
		})
	}

	user, err := repository.GetUserByEmail(body.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	newUser, err := repository.GetUserById(user.ID.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(newUser.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email and password mismatch",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":      newUser.ID,
		"expires": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
