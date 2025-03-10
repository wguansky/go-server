package handlers

import (
	"context"
	"net/http"
	"time"
	"workflow/configs"
	"workflow/models"
	"workflow/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	var json struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := models.CreateUser(json.Username, json.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func LoginHandler(c *gin.Context) {
	var json struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// JWT validate
	if user, err := models.GetUserByUsername(json.Username); err == nil && user.Password == json.Password {
		token, err := utils.GenerateJWT(user.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT"})
			return
		}

		// 存储会话到 Redis
		err = configs.Redis.Set(context.Background(), "session:"+json.Username, token, 24*time.Hour).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store session"})
			return
		}

		// Session Login
		session := sessions.Default(c)
		session.Set("user_id", user.ID)
		session.Set("username", user.Username)
		session.Save()

		c.JSON(http.StatusOK, gin.H{"token": token})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
}

func LogoutHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	var json struct {
		Username string `json:"username"`
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := configs.Redis.Del(context.Background(), "session:"+json.Username).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

func ProtectedHandler(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Hello, authenticated user"})
}
