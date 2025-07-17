package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		fmt.Printf("[%s] %s %s %v %d",
			c.Request.Method, c.Request.URL, c.Request.Host, c.Writer.Header(), duration)
	}
}
