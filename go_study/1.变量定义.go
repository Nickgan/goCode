package main

import (
	"fmt"
	"goCode/go_study/version"
)

var (
	name = "全局变量name 张三"
	age  = 18
)

func main() {

	fmt.Println("hello world")

	// 1.先定义，在复制
	var name string
	name = "张三"
	fmt.Println("name==========>", name)

	// 2.声明同时赋值，这里可以省略数据类型
	var age int = 18
	fmt.Println("age==========>", age)

	// 3.声明并赋值，又叫段声明
	address := "北京"
	fmt.Println("address==========>", address)

	fmt.Println("name==========>", name)

	const 常量 = "常量,要用const定义，定义了就不能改了"
	fmt.Println("name3==========>", 常量)

	fmt.Println("=========================")

	fmt.Printf("引用外面的变量: %v", version.Version2)
}
