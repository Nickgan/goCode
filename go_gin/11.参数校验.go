package main

import (
	"fmt"
	"goCode/go_gin/resp"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
*
* 参考文档： https://www.fengfengzhidao.com/special/3/41
 */

func main() {

	gin.SetMode(gin.ReleaseMode) //精简日志输出
	r := gin.Default()

	// 4：json参数(结构体属性名要大写)
	r.POST("/user/add/json/validata", func(c *gin.Context) {
		type User struct {
			Name string `json:"name" binding:"required,min=3,max=5"`
			Age  int    `json:"age" binding:"gte=10,lte=15"`
			Sex  string `json:"sex" binding:"required,oneof=男 女"`
		}

		var user User
		err := c.ShouldBindJSON(&user)
		if err != nil {
			fmt.Println("error============>", err.Error())
			// 校验失败，返回 400 错误
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(user)
		resp.OkWithMsg(c, "操作成功")
	})

	fmt.Println("启动成功...")
	r.Run(":8080")
}
