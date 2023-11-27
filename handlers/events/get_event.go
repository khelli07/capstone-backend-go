package events

import (
	"backend-go/ds"
	"backend-go/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetEventById(c *gin.Context) {
	s := c.Param("id")
	id, err := strconv.Atoi(s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	entity, err := repository.GetEventById(ds.CTX, ds.DS, int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, entity)
}
