package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	c := gin.Default()

	c.GET("/", func(c *gin.Context) {

		name := c.Query("name")
		age := c.DefaultQuery("age", "18")
		keyList := c.QueryArray("key")
		fmt.Println("name:", name, "age:", age, "keyList:", keyList)
	})

	c.Run(":8080")
}
