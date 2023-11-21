package main

import (
	"backend-go/database"
	handlers "backend-go/handlers/user"
	"backend-go/middlewares"
	"backend-go/models"
	"backend-go/utils"
	"net/http"
	"os"

	_ "backend-go/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	utils.LoadEnv()
	database.Connect()
	database.Migrate()
}

// @title           Match-Event Backend API
// @version         1.0
// @description     A matchmaking service API in Go using Gin framework.

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:9999
// @BasePath  /
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // http://[HOST_URL]/swagger/index.html

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
