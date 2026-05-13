package main

import "fmt"

func f1[T int | float64](a, b T) {
	fmt.Println((a + b))
}

func main() {
	f1(12.55, 22.11)
}
