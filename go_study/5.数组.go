package main

import "fmt"

func main() {

	var arr = [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Println(arr[0])

	for i := range arr {
		fmt.Println(arr[i])
	}

	fmt.Println("最后一个元素===>", arr[len(arr)-1])

	var arr2 = [...]string{"a", "b", "c", "d", "e"}
	fmt.Println(arr2)
	fmt.Println("最后一个元素===>", arr2[len(arr2)-1])

}
