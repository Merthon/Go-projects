package middlewares

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// AuthRequired 演示简单鉴权，检查 Header 中的 Authorization
func AuthRequired() gin.HandlerFunc {
	return func (c *gin.Context) {
		token := c.GetHeader("Authorization")
		//简单演示
		if token != "Bearer secrettoken" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return 
		}
		c.Next()
	}
}