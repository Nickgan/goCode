package main

import "fmt"

func main() {
	var name string

	fmt.Print("请输入你的名字：")
	n, error := fmt.Scan(&name)
	fmt.Println("你的名字是：", name)

	fmt.Println("n:", n)
	fmt.Println("error:", error)
}
