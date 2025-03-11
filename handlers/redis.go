package handlers

import (
	"context"
	"net/http"
	"rest-api/configs"

	"github.com/gin-gonic/gin"
)

func RedisHandler(c *gin.Context) {
	err := configs.Redis.Set(context.Background(), "key", "Hello from Redis!", 0).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	val, err := configs.Redis.Get(context.Background(), "key").Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": val})
}
