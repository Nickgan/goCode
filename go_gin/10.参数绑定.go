package main

import (
	"fmt"
	"goCode/go_gin/resp"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode) //精简日志输出
	r := gin.Default()

	//1:绑定查询参数
	r.GET("/index", func(c *gin.Context) {

		type User struct {
			Name string `form:"name"`
			Age  int    `form:"age"`
		}

		var user User
		c.BindQuery(&user)

		fmt.Println("bindQuery, user==========>", user)

		// 返回json
		resp.OkWithMsg(c, "操作成功")
	})

	// 2：路径参数
	r.GET("/user/:name/:id", func(c *gin.Context) {
		type User struct {
			Name string `uri:"name"`
			Id   int    `uri:"id"`
		}

		var user User
		c.ShouldBindUri(&user)
		fmt.Println(user)

		// 返回json
		resp.OkWithMsg(c, "操作成功")
	})

	// 3：表单参数
	r.POST("/user/add", func(c *gin.Context) {
		type User struct {
			Name string `form:"name"`
			Age  int    `form:"age"`
		}

		var user User
		c.ShouldBind(&user)
		fmt.Println(user)

		// 返回json
		resp.OkWithMsg(c, "操作成功")
	})

	// 4：json参数
	r.POST("/user/add/json", func(c *gin.Context) {
		type User struct {
			Name string `json:"name" binding:"required,min=3,max=5"`
			Age  int    `json:"age"`
		}

		var user User
		c.ShouldBindJSON(&user)
		fmt.Println(user)
		resp.OkWithMsg(c, "操作成功")
	})

	// 5：header参数
	r.POST("/headerparam/demo", func(c *gin.Context) {
		type User struct {
			Name        string `header:"Name"`
			Age         int    `header:"Age"`
			UserAgent   string `header:"User-Agent"`
			ContentType string `header:"ontent-Type"`
		}

		fmt.Println("hedaer")
		var user User
		err := c.ShouldBindHeader(&user)
		fmt.Println(user, err)
		fmt.Println("=====>", user.ContentType)
		resp.OkWithMsg(c, "操作成功")
	})

	r.Run(":8080")
}
