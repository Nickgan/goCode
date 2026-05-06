package main

import "fmt"

func main() {

	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println(arr)

	var arr2 = [3]int{1, 2, 5}

	arr2[2] = 10
	fmt.Println(arr2)

	fmt.Println(len(arr2))

	//声明一个切片slice
	var s1 []string
	s1 = make([]string, 0)
	s1 = make([]string, 0, 10)

	fmt.Println(s1 == nil)
}
