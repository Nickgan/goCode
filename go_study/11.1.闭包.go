package main

import "fmt"

// makeCounter 是一个“工厂函数”，它返回一个闭包
func makeCounter() func() int {

	count := 0 // 这是被捕获的“自由变量”
	return func() int {
		count++      // 闭包内部可以修改外部变量
		return count // 并返回修改后的值
	}
}

func main() {

	// 调用 makeCounter，得到一个闭包，赋值给 counter
	counter := makeCounter()

	// 每次调用 counter()，它都会记住上一次的 count 值
	fmt.Println(counter()) // 输出: 1
	fmt.Println(counter()) // 输出: 2
	fmt.Println(counter()) // 输出: 3

	// 创建另一个独立的计数器
	anotherCounter := makeCounter()
	fmt.Println(anotherCounter()) // 输出: 1 (拥有自己独立的 count 变量)

}
