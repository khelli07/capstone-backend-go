package middlewares

import "github.com/gin-gonic/gin"

func AdminOnly(c *gin.Context) {
	c.Next()
}
