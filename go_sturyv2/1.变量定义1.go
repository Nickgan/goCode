package main

import "fmt"

func main() {

	// 1.先定义后赋值
	var name string
	name = "张三"
	fmt.Println("name==========>", name)

	// 2.声明并赋值
	var age = 18
	fmt.Println("age==========>", age)

	// 3.短声明
	address := "北京"
	fmt.Println("address==========>", address)

	// 4.定义多个变量
	var name1, name2 string
	name1, name2 = "张三", "张三"
	fmt.Println("name1==========>", name1, "name2==========>", name2)
}
