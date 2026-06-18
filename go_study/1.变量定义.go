package main

import (
	"fmt"
	"goCode/go_study/version"
)

var (
	name = "全局变量name 张三"
	age  = 18
	sex  = true
)

// 常量定义（定义的时候就要赋值，并且不支持修改了）
const version111 = "10.1"

func main() {

	fmt.Println("hello world")

	// 1.先声明，在复制
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

	var f1 = -10.5
	fmt.Println(f1)

	fmt.Printf("%v,%T", f1, f1)

	var 中文字符 rune = '中'
	fmt.Printf("================>%c \n", 中文字符)

	fmt.Printf("%v", version111)

	fmt.Println(sex)

}

// 定义多个变量
func 多个变量() {

	var a, b int
	var a1, b1 = 1, 2
	a2, b2 := 2, 3
	fmt.Println(a, b, a1, b2, b1, a2)
}
