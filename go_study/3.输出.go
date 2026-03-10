package main

import "fmt"

func main() {
	fmt.Println("请输入你的名字:")
	var name string
	fmt.Scan(&name)
	fmt.Println("你输入的名字是:", name)

	fmt.Println("请输入你的年龄:")
	var age int
	n, error := fmt.Scan(&age)
	if error != nil {
		fmt.Println("输入错误", error)
		fmt.Println(n)
		return
	}
	fmt.Println("你输入的年龄是:", age)
}
