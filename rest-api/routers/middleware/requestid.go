package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check Header được gửi đến, sử dụng nó nếu nó tồn tại
		requestId := c.Request.Header.Get("X-Request-Id")

		// Tạo request id với UUID4
		if requestId == "" {
			u4, _ := uuid.NewV4()
			requestId = u4.String()
		}

		// Export nó để sử dụng trong ứng dụng
		c.Set("X-Request-Id", requestId)

		// Set X-Request-Id cho header
		c.Writer.Header().Set("X-Request-Id", requestId)
		c.Next()
	}
}
