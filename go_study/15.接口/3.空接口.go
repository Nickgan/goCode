package main

import "fmt"

//1：空接口可以接受任何类型
//2：换一种，任何类型都实现了空接口的定义

type NullInterface interface {
}

// MyPrint 空接口可以接受任何一种类型（第一种写法，直接写一个空接口）
//func MyPrint(a NullInterface) {
//	fmt.Println(a)
//}

// MyPrint 空接口可以接受任何一种类型（第二种写法，写一个interface{}）
func MyPrint(val interface{}) {
	fmt.Println(val)
}

// MyPrint 空接口可以接受任何一种类型（第三种写法，写一个any）
func MyPrint1(val any) {
	fmt.Println(val)
}

type MyCat struct {
	Name string
	Age  int
}

func main() {
	MyPrint(MyCat{Name: "小猫", Age: 1})
	MyPrint(1)
}
