package handlers

import (
	"net/http"
	"rest-api/configs"

	"github.com/gin-gonic/gin"
)

func DBHandler(c *gin.Context) {
	var result string
	err := configs.DB.QueryRow("SELECT 'Hello from PostgreSQL!'").Scan(&result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": result})
}
