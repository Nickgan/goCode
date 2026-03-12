package main

import (
	"fmt"
	"goCode/sgg_study/utils"
)

func main() {
	fmt.Println("hello world")

	fmt.Println(ff1())

	fmt.Println(utils.Cal(1, 2, "a"))
}

func ff1() int {
	fmt.Println("ff1")
	return 1
}
