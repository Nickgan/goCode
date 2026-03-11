package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 10

	fmt.Println(&a)

	var bpoint *int = &a
	fmt.Println("bpoint = ", bpoint)

	fmt.Println("*bpoint = ", bpoint)

	*bpoint = 20
	fmt.Println("a = ", a)

	n1, _ := test()
	fmt.Println("test() = ", n1)

	test2()

	//test3()
}

func test3() {

	i := 10
	i--
	var b int = 20 - i

	fmt.Println("=================b = ", b)
}

func test() (bool, bool) {
	fmt.Println("test.....")
	return true, false
}

func test2() {
	fmt.Println("---------------------------")

	a := 10
	b := 20
	fmt.Printf("交换前的值是: a = %v, b = %v", a, b)

	var str string = "123aa"
	n1, er1 := strconv.ParseInt(str, 10, 64)

	if er1 != nil {
		fmt.Println("转换失败")
	}

	fmt.Println("n1 = ", n1, " er1 = ", er1)

	//t := a
	//a = b
	//b = t

	//a = a + b
	//b = a - b
	//a = a - b

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Printf("交换后的值是: a = %v, b = %v", a, b)

}
