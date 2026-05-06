package main

import "fmt"

func main() {
	fmt.Println("hello world")

	list1 := make([]int, 2, 2)
	fmt.Println(list1, len(list1), cap(list1))

}
