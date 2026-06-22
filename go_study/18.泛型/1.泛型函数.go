package main

import "fmt"

// 1：单个类型的泛型函数
func f1[T int | float64](a, b T) {
	fmt.Println((a + b))
}

// 2：多个类型的泛型函数
func f11[T int | int32, K string | float64](a T, b K) {
	fmt.Println(a, b)
}

// 3: 1定义一个接口包含所有数字类型
type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr
}

// 3: 2使用上面定义的接口，泛型函数
func f2[T Number](a, b T) {
	fmt.Println(a + b)
}

func main() {
	f1(12.55, 22.11)
}
