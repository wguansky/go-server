package middleware

import (
	"workflow/configs"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/postgres"
	"github.com/gin-gonic/gin"
)

func SessionMiddleware() gin.HandlerFunc {
	// store:= cookie.NewStore([]byte(viper.GetString("session_secret")))
	store, _ := postgres.NewStore(configs.DB)
	return sessions.Sessions("session", store)
}
