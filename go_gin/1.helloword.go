package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// 1. 创建默认的 Gin 引擎（自带了日志 Logger 和崩溃恢复 Recovery 中间件）
	gin.SetMode(gin.ReleaseMode) //精简日志输出
	r := gin.Default()

	// 2. 定义一个 GET 请求的路由
	r.GET("/", func(c *gin.Context) {
		// 3. 返回 JSON 格式的响应
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Gin!",
			"status":  200,
		})

	})

	// 4. 启动服务器，监听 8080 端口
	r.Run(":8080")
}
