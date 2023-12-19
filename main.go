package main

import (
	"backend-go/handlers/categories"
	"backend-go/handlers/events"
	"backend-go/handlers/ml"
	"backend-go/handlers/reviews"
	"backend-go/handlers/users"
	"backend-go/middlewares"
	"backend-go/mongodb"
	"backend-go/utils"
	"net/http"
	"os"

	_ "cloud.google.com/go/datastore"

	_ "backend-go/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	utils.LoadEnv()
	utils.InitGCS()
	mongodb.ConnectDB()
	mongodb.InitCollections()
}

// @title           Match-Event Backend API
// @version         1.0
// @description     A matchmaking service API in Go using Gin framework.

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath  /
func main() {
	defer mongodb.Client.Disconnect(mongodb.Context)

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
		auth.POST("/register", users.Register)
		auth.POST("/login", users.Login)
	}

	user := router.Group("/users")
	{
		user.GET("/", middlewares.RequireAuth, users.GetUserInfo)
		user.PUT("/", middlewares.RequireAuth, users.UpdateUser)
		user.GET("/:id", users.GetUserById)
	}

	event := router.Group("/events")
	{
		event.GET("/", events.GetEvents)
		event.GET("/:id", events.GetEventById)
		event.GET("/popular", ml.GetPopularEvents)
		event.GET("/reco", middlewares.RequireAuth, ml.GetRecommendedEvents)
		event.POST("/", events.CreateEvent)
		event.PUT("/:id", events.UpdateEvent)
		event.DELETE("/:id", events.DeleteEvent)
		event.POST("/:id/join", middlewares.RequireAuth, events.JoinEvent)
		event.POST("/:id/image", events.UploadEventImage)
	}

	review := router.Group("/reviews")
	{
		review.GET("/:event_id", reviews.GetReviews)
		review.POST("/:event_id", middlewares.RequireAuth, reviews.CreateReview)
		review.DELETE("/:id", middlewares.RequireAuth, reviews.DeleteReview)
	}

	category := router.Group("/categories")
	{
		category.GET("/", categories.GetCategories)
		category.POST("/", categories.CreateCategory)
		category.PUT("/:id", categories.UpdateCategory)
		category.DELETE("/:id", categories.DeleteCategory)
	}

	router.Run(":" + os.Getenv("PORT"))
}
