package main

import "fmt"

var a = test1()

func test1() int {
	fmt.Println("test 1()....")
	return 80
}

func init() {

	fmt.Println("init ()...===a====>", a)
	a = 100
	fmt.Println("init ()...")
}

func main() {
	fmt.Println("main ()=====>", a)
}
