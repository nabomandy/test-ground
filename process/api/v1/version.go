package v1

import (
	"github.com/gin-gonic/gin"
)

func Version(c *gin.Context) {
	c.JSON(200, gin.H {
		"version": "1.0.0",
	})
}