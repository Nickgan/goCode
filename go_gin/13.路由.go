package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 基础路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "welcom"})
	})

	// 单个路径参数
	r.GET("/user/:id", func(c *gin.Context) {
		var id = c.Param("id")
		c.JSON(http.StatusOK, gin.H{"message": id})
	})

	// 多个路径参数
	r.GET("/user/:id/:name", func(c *gin.Context) {
		var id = c.Param("id")
		var name = c.Param("name")
		c.JSON(http.StatusOK, gin.H{"id": id, "name": name})
	})

	// 通配符参数
	// 使用场景：
	//   - 文件路径：/files/*filepath 可以匹配 /files/docs/readme.md、/files/images/photo.jpg 等
	//   - 静态资源：/static/*filepath 可以匹配所有静态资源路径
	//   - 代理转发：将剩余路径转发到其他服务
	//
	r.GET("/files/*filePath", func(c *gin.Context) {
		var filePath = c.Param("filePath")
		c.JSON(http.StatusOK, gin.H{"message": filePath})
	})

	type UserQueryReq struct {
		Name string `form:"name"  binding:"required"`
		Age  int    `form:"age" binding:"required,gte=0,lte=120"`
	}

	// 查询参数
	r.GET("/search", func(c *gin.Context) {
		//var name = c.Query("name")
		//var age = c.Query("age")
		//c.JSON(http.StatusOK, gin.H{"name": name, "age": age})
		var userQueryReq UserQueryReq
		//if err := c.ShouldBind(&userQueryReq); err != nil {
		if err := c.ShouldBindQuery(&userQueryReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"name": userQueryReq.Name, "age": userQueryReq.Age})
	})

	// 表单post参数
	r.POST("/form", func(c *gin.Context) {
		var name = c.PostForm("name")
		var age = c.PostForm("age")
		c.JSON(http.StatusOK, gin.H{"name": name, "age": age})
	})

	// JSON post参数
	type UserReq struct {
		Name  string `json:"name"  binding:"required"`
		Email string `json:"email" binding:"required,email"`
		Age   int    `json:"age" binding:"required,gte=0,lte=120"`
	}

	r.POST("/api/users11", func(c *gin.Context) {
		var userReq UserReq
		//if err := c.ShouldBind(&userReq); err != nil {
		if err := c.ShouldBindJSON(&userReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"name": userReq.Name, "age": userReq.Age, "email": userReq.Email})
	})

	// 路由分组
	v1 := r.Group("/api/v1")
	{
		v1.GET("/users", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "get users"})
		})
		v1.POST("/users", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "post users"})
		})
	}

	// 嵌套分组
	api := r.Group("/api2")
	{
		var userApiV1 = api.Group("/v1/user2")
		{
			userApiV1.GET("/users", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "get users"})
			})
			userApiV1.POST("/users", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "post users"})
			})
		}

		var userAPI2 = api.Group("/v2/user")
		{
			userAPI2.GET("/users", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "get users"})
			})
			userAPI2.POST("/users", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "post users"})
			})
		}
	}

	//any
	r.Any("/any", func(c *gin.Context) {
		var method = c.Request.Method
		c.JSON(http.StatusOK, gin.H{"method": method})
	})

	// 启动服务器
	r.Run(":80")
	fmt.Println("服务器启动成功.")
}
