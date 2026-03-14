package main

func main() {

	//1： 将匿名函数赋值给一个变量
	f1 := func(n1 int) {
		println("hello world============>", n1)
	}

	f1(100)

	//2：声明的同时调用匿名函数
	sum := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 30)

	println(sum)
}
