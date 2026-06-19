package main

import "fmt"

func main() {

	//sum1 := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 55)
	//fmt.Println("main sum==========>", sum1)

	// 匿名函数
	//m2 := func() {
	//	fmt.Println("匿名函数被调用==========>")
	//}
	//m2()
	//
	//ff()

	var n int = 10
	fmt.Println("main n=", &n)
	ts1(&n)
	fmt.Println("main n=", n)

	// %p 打印地址
	var list = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("list=%p", &list) // %p打印地址

}

func sum(numList ...int) (sum int) {

	fmt.Printf("numList==========> %T \n", numList)

	for n, i := range numList {
		sum += i
		fmt.Println("n=", n, "i=", i)
	}
	fmt.Println("sum=", sum)
	return
}

func ff() {
	var f int
	fmt.Println("请输入一个数字：")
	fmt.Scan(&f)

	var mapFunc = map[int]func(){
		1: func() {
			fmt.Println("1 调用")
		},
		2: func() {
			fmt.Println("2 调用")
		},
	}

	mapFunc[f]()

	//fmt.Println(mapFunc)

}

// &是取地址，*是解引用，去这个地址指向的值
func ts1(n *int) {
	fmt.Println("ts1 n=", n)
	fmt.Println("ts1 n=", *n)
	*n = 18
	fmt.Println("ts1 n=", *n)
}

// 接收一个函数，返回一个函数
func f111111(a func() (int, string)) (func(), int) {
	a()
	return func() {
		fmt.Println("f111111")
	}, 1
}
