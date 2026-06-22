package main

import (
	"fmt"
)

func read() {

	//defer func() {
	//	if err := recover(); err != nil {
	//		//fmt.Println(err)
	//		fmt.Println("发生异常了，异常信息=====>", err)
	//	}
	//}()

	var list = []int{1, 2, 3, 4, 5}
	fmt.Println(list[10])
}

func main() {

	defer func() {

		if err := recover(); err != nil {
			//fmt.Println(err)
			fmt.Println("发生异常了，异常信息=====>", err)

			//打印堆栈信息(让知道哪里错了)
			//fmt.Println(string(debug.Stack()))
		}
	}()

	read()

	fmt.Println("main方法结束。。。")
}
