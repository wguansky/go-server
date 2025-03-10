package middleware

import (
	"workflow/configs"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/postgres"
	"github.com/gin-gonic/gin"
)

func SessionMiddleware() gin.HandlerFunc {
	// store:= cookie.NewStore([]byte(viper.GetString("session_secret")))
	store, err := postgres.NewStore(configs.DB, []byte("secret-key"))
	if err != nil {
		panic("Can't connect to sesssion database!")
	}
	return sessions.Sessions("session", store)
}
