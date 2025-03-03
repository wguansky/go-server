package routes

import (
	"workflow/handlers"
	"workflow/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/register", handlers.RegisterHandler)
	r.POST("/login", handlers.LoginHandler)
	r.GET("/logout", handlers.LogoutHandler)

	auth := r.Group("/")
	auth.Use(middleware.SessionMiddleware())
	{
		auth.GET("/hello", handlers.HelloHandler)
		auth.GET("/db", handlers.DBHandler)
		auth.GET("/redis", handlers.RedisHandler)
		auth.GET("/protected", handlers.ProtectedHandler)
	}
}
