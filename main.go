package main

import (
	"rest-api/middleware"
	"rest-api/routes"

	"github.com/gin-gonic/gin"

	"rest-api/configs"
)

func main() {
	defer configs.DB.Close()
	r := gin.Default()
	r.Use(middleware.SessionMiddleware())
	routes.RegisterRoutes(r)
	r.Run("0.0.0.0:8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
