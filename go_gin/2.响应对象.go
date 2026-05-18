package main

import (
	"goCode/go_gin/resp"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	resp.OkWithMsg(c, "操作成功")
}

func main() {

	gin.SetMode(gin.ReleaseMode) // 精简日志输出

	// 1: 初始化引擎
	r := gin.Default()

	// 2： 挂载处理路由
	r.GET("/index", Index)

	// 3： 绑定端口，启动
	r.Run(":8080")

}
