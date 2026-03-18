package main

import "fmt"

// AliasCode 定义一个类型别名
type AliasCode = int

// MyCode 定义一个自定义类型
type MyCode int

var (
	SuccessCode      MyCode    = 100
	SuccessAliasCode AliasCode = 100
)

// 自定义类型可以绑定方法
func (c MyCode) getCode() {
	fmt.Println(c)
}

// 类型别名不可以绑定方法
//func (c AliasCode) getCode() {
//	fmt.Println(c)
//}

func main() {
	/**
	 * 1:自定义类型别名不能绑定方法
	 * 2:类型别名和原类型可以直接比较
	 * 3:类型别名打印类型还是原类型
	 */

	fmt.Printf("%T, %T \n", SuccessCode, SuccessAliasCode)

	// 类型别名和原类型可以直接比较
	var i = 100
	fmt.Println(SuccessAliasCode == i)

	//自定义类型不能跟变量进行比较，要转换为原类型才能和原类型进行比较
	fmt.Println(int(SuccessCode) == i)
}
