package main

import "fmt"

func main() {
	var a int = 10

	fmt.Println(&a)

	var bpoint *int = &a
	fmt.Println("bpoint = ", bpoint)

	fmt.Println("*bpoint = ", bpoint)

	*bpoint = 20
	fmt.Println("a = ", a)
}
