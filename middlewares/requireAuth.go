package middlewares

import (
	"backend-go/fs"
	"backend-go/models"
	"backend-go/repository"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "No token provided",
		})
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	tokenString = strings.Split(tokenString, "Bearer ")[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["expires"].(float64) {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Token expired",
			})
			c.AbortWithStatus(http.StatusBadRequest)
		}

		user, err := repository.GetUserById(fs.CTX, fs.FSClient, claims["id"].(string))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": err.Error()})
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var tokenUser models.TokenUser
		tokenUser = models.TokenUser{
			ID:       claims["id"].(string),
			Username: user.Username,
			Email:    user.Email,
		}

		c.Set("user", tokenUser)
		c.Next()
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid token",
		})
		c.AbortWithStatus(http.StatusBadRequest)
	}

}
