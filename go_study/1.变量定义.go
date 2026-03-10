package main

import (
	"fmt"
	"goCode/go_study/version"
)

var age int

var (
	age1 int = 18
	age2 int = 19
)

const qName string = "测试常量定义，定义的时候就要赋值，赋值了就不能改了"

func main() {

	// 先定义再赋值
	var name string
	name = "张三"
	fmt.Println(name)

	// 声明变量后就赋值
	var userName string = "枫枫"
	fmt.Println(userName)

	var name1, name2, name3 string = "张三", "李四", "王五"
	fmt.Println(name1, name2, name3)

	a1, a2 := 100, 200
	fmt.Println(a1, a2)

	fmt.Println(age2)
	fmt.Println(qName)

	// 输出引用文件的常量
	fmt.Println(version.Versin1)
}
