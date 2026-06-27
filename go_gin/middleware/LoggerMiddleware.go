package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// ========== 日志中间件 ==========
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 前置处理
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		// 进入下一个处理函数
		c.Next()

		// 后置处理
		latency := time.Since(start)
		status := c.Writer.Status()
		fmt.Printf("[%s] %s %d %v\n", method, path, status, latency)
	}
}
