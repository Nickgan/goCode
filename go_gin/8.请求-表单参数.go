package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	c := gin.Default()

	c.POST("/index/add", func(c *gin.Context) {

		age, ok := c.GetPostForm("age")
		name := c.PostForm("name")

		fmt.Println("name:", name, "age:", age, "ok:", ok)
	})

	c.Run(":8080")
}
