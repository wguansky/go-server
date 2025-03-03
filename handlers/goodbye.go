package handlers

import (
	"github.com/gin-gonic/gin"
)

func GoodbyeHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Goodbye, World!",
	})
}
