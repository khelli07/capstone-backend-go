package main

import (
	"backend-go/database"
	handlers "backend-go/handlers/user"
	"backend-go/middlewares"
	"backend-go/models"
	"backend-go/utils"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	utils.LoadEnv()
	database.Connect()
	database.Migrate()
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API is up and running!",
		})
	})

	auth := router.Group("/auth")
	{
		auth.GET("/", middlewares.RequireAuth, func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "You are authenticated!",
				"user":    c.MustGet("user").(models.PublicUser),
			})
		})
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	}

	router.Run(":" + os.Getenv("PORT"))
}
