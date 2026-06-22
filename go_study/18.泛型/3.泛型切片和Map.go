package main

import "fmt"

// 定义泛型切片和Map
type MyList[T any] []T

type MyMap[K string | int, V any] map[K]V

func main() {

	var list MyList[int]
	list = append(list, 1, 2, 3, 4, 5)
	fmt.Println(list)

	var mp = make(MyMap[string, int])
	mp["a"] = 1
	mp["b"] = 2
	fmt.Println(mp)
}
