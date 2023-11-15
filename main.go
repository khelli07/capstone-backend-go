package main

import (
	"backend-go/database"
	"backend-go/utils"
	"net/http"

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

	router.Run(":9999")
}
