package main

import "github.com/gin-gonic/gin"

func main() {
	gin.SetMode(gin.ReleaseMode) // 精简日志输出

	c := gin.Default()

	// 前面是访问别名（一般就叫static，为了方便理解，这里写作st），后面是路径，静态文件的路径不能再被路由使用了
	c.Static("st", "static")

	c.Run(":8080")
}
