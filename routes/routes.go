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

	// 受会话保护的路由
	sessionProtected := r.Group("/")
	sessionProtected.Use(middleware.SessionMiddleware())
	{
		sessionProtected.GET("/session/hello", handlers.HelloHandler)
		sessionProtected.GET("/session/db", handlers.DBHandler)
		sessionProtected.GET("/session/redis", handlers.RedisHandler)
		sessionProtected.GET("/session/protected", handlers.ProtectedHandler)
	}

	// 受 JWT 保护的路由
	jwtProtected := r.Group("/")
	jwtProtected.Use(middleware.JWTAuthMiddleware())
	{
		jwtProtected.GET("/jwt/hello", handlers.HelloHandler)
		jwtProtected.GET("/jwt/db", handlers.DBHandler)
		jwtProtected.GET("/jwt/redis", handlers.RedisHandler)
		jwtProtected.GET("/jwt/protected", handlers.ProtectedHandler)
	}
}
