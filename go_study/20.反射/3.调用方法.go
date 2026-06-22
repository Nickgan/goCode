package main

import (
	"fmt"
	"reflect"
)

func add(a int, b int) int {
	return a + b
}

type U struct {
	Name string
}

func (u U) Hello(msg string) string {
	return u.Name + " say: " + msg
}

func main() {

	// 1: 调用普通方法
	fn := reflect.ValueOf(add)

	// 构造参数
	args := []reflect.Value{reflect.ValueOf(1), reflect.ValueOf(2)}

	// 调用
	res := fn.Call(args)
	fmt.Println(res[0].Int())

	//=====================================
	// 2:动态调用结构体方法

	u := U{
		Name: "张三",
	}

	t1 := reflect.ValueOf(&u)
	method := t1.MethodByName("Hello")

	args1 := []reflect.Value{reflect.ValueOf("hello world")}

	res2 := method.Call(args1)
	fmt.Println(res2[0].String())

}
