package reviews

import (
	"backend-go/models"
	"backend-go/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateReview(c *gin.Context) {
	var body struct {
		Rating  int    `json:"rating"`
		Comment string `json:"comment"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if body.Rating < 1 || body.Rating > 5 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "rating must be between 1 and 5"})
		return
	}

	review := models.Review{
		EventID: c.Param("event_id"),
		UserID:  c.MustGet("user").(models.TokenUser).ID,
		Rating:  body.Rating,
		Comment: body.Comment,
	}

	result, err := repository.CreateReview(&review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": result.InsertedID})
}
