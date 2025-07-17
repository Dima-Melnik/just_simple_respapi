package middleware

import (
	"just_simple_api/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSONMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "POST" || c.Request.Method == "PUT" {
			if c.Request.Header.Get("Content-Type") != "application/json" {
				utils.SendError(c, http.StatusUnsupportedMediaType, "Content-Type must be application/json")
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
