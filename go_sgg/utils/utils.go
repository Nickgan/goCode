package utils

import "fmt"

func Cal(a int, b int, op string) int {

	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		fmt.Println("不支持该运算符")
		return 0
	}
}
