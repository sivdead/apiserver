package middleware

import (
	"time"
	
	"github.com/gin-gonic/gin"
	. "github.com/sivdead/apiserver/handler"
	"github.com/sivdead/apiserver/pkg/errno"
	"github.com/sivdead/apiserver/pkg/token"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		if c.Request.URL.Path == "/login" {
			return
		}
		
		// Parse the json web token.
		if context, err := token.ParseRequest(c); err != nil || context.EXP < time.Now().Unix() {
			SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}
		
		c.Next()
	}
}
