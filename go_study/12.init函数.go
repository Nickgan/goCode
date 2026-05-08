package main

import (
	"fmt"
)

var m12 int = m111()

func init() {
	fmt.Println("init ()...111")
}

func init() {
	fmt.Println("init ()...222")
}

func init() {
	fmt.Println("init ()...333")
}

func m111() int {
	fmt.Println("=====================m11()")
	return 123444
}

func main() {
	fmt.Println("main ()...")
}
