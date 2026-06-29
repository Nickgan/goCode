package middleware

import (
	"goCode/go_blog/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 Header 获取 Token
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(401, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// 检查Bearer前缀
		tokenString := strings.TrimPrefix(token, "Bearer ")
		if tokenString == token {
			utils.Unauthorized(c, "Bearer token is required")
			c.Abort()
			return
		}

		// 解析token
		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			utils.Unauthorized(c, "Invalid token")
			c.Abort()
			return
		}

		// 将用户信息存储到上下文中
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Next()
		c.Next()
	}
}
