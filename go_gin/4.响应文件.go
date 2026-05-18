package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/download", func(c *gin.Context) {

		fmt.Println("--------")
		c.Header("Content-Type", "application/octet-stream")              // 表示是文件流，唤起浏览器下载，一般设置了这个，就要设置文件名
		c.Header("Content-Disposition", "attachment; filename=4.响应文件.go") // 用来指定下载下来的文件名
		c.File("go_gin/4.响应文件.go")
	})

	r.Run(":8080")
}
