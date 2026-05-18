package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	c := gin.Default()

	c.GET("/index/:id", func(c *gin.Context) {
		id := c.Param("id")
		fmt.Println("id:", id)
	})

	c.Run(":8080")
}
