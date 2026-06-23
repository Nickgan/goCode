package main

import (
	"fmt"
	"os"
)

func globalRecorver() {
	if r := recover(); r != nil {
		fmt.Println("发生异常了，异常信息=====>", r)
	}
}

func read2(finaName string) {
	// 一次性读取文件
	byteData, err := os.ReadFile(finaName)

	if err != nil {
		fmt.Println("文件读取失败")
		panic(err)
	}

	fmt.Println(string(byteData))
}

func main() {
	defer globalRecorver()
	read2("go_study/17.异常处理/3.恢复使用2.go")
	fmt.Println("mai方法结束。。。")
}
