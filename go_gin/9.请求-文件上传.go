package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	c := gin.Default()

	c.POST("/users", func(c *gin.Context) {

		fileHeader, err := c.FormFile("file")
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(fileHeader.Filename) //文件名
		fmt.Println(fileHeader.Size)     //文件大小

		file, _ := fileHeader.Open()
		byteData, _ := io.ReadAll(file)

		err = os.WriteFile("xxx.jpg", byteData, 0666)
		fmt.Println(err)
	})

	c.Run(":8080")
}
