// @Title
// @Description
// @Author
// @Update
package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func SessionMiddleware() gin.HandlerFunc {
	store := cookie.NewStore([]byte(viper.GetString("session_secret")))
	return sessions.Sessions("mysession", store)
}
