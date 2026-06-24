package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {

	//绝对路径：   D:\codes\goCode\go_study\19.文件操作\hello.txt

	// 一次性读取文件
	byteData, err := os.ReadFile("go_study/19.文件操作/hello.txt")

	if err != nil {
		fmt.Println("文件读取失败", err)
		return
	}

	fmt.Println(string(byteData))

	// 获取当前文件信息
	_, filePath, l, ok := runtime.Caller(0)
	fmt.Println(filePath, l, ok)

}
