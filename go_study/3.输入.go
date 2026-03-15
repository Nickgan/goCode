package main

import "fmt"

func main() {
	var name string

	fmt.Print("请输入你的名字：")
	fmt.Scan(&name)
	fmt.Println("你的名字是：", name)
}
