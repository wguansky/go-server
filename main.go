package main

import (
	"workflow/middleware"
	"workflow/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// Use the session middleware
	r.Use(middleware.SessionMiddleware())

	routes.RegisterRoutes(r)

	r.Run("0.0.0.0:8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
