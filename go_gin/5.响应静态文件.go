package main

import "github.com/gin-gonic/gin"

func main() {
	gin.SetMode(gin.ReleaseMode) // 精简日志输出

	c := gin.Default()

	// 前面是访问别名，后面是路径
	c.Static("st", "static")

	c.Run(":8080")
}
