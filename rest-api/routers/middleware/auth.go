package middleware

import (
	"rest-api/controllers"
	"rest-api/pkg/errno"
	"rest-api/pkg/token"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		if _, err := token.ParseRequest(c); err != nil {
			controllers.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
