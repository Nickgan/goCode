package main

import (
	"fmt"
	"goCode/go_gin/resp"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode) //精简日志输出
	r := gin.Default()

	// 4：json参数
	r.POST("/user/add/json", func(c *gin.Context) {
		type User struct {
			Name string `json:"name" binding:"required,min=3,max=5"`
			Age  int    `json:"age"`
		}

		var user User
		err := c.ShouldBindJSON(&user)
		fmt.Println("error============>", err.Error())
		if err != nil {
			// 校验失败，返回 400 错误
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(user)
		resp.OkWithMsg(c, "操作成功")
	})

	r.Run(":8080")
}
