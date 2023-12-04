package users

import (
	"backend-go/fs"
	"backend-go/models"
	"backend-go/repository"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var body struct {
		Username  string
		Email     string
		Password  string
		Password2 string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to process request body",
		})
		return
	}

	if body.Password != body.Password2 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password mismatch",
		})
		return
	}

	doc, err := repository.GetUserByEmail(fs.CTX, fs.FSClient, body.Email)
	fmt.Println(doc)
	if doc != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email has been taken!",
		})
		return
	} else if err != nil && err.Error() != "User not found" {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	user := models.User{
		Username: body.Username,
		Email:    body.Email,
		Password: string(hash),
	}
	_, err = repository.CreateUser(fs.CTX, fs.FSClient, &user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
	})
}
