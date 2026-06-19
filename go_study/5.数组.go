package main

import "fmt"

func main() {

	var arr = [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Println(arr[0])
	fmt.Printf("%T \n", arr)

	for i := range arr {
		fmt.Println(arr[i])
	}

	fmt.Println("最后一个元素===>", arr[len(arr)-1])

	var arr2 = [...]string{"a", "b", "c", "d", "e"}
	fmt.Println(arr2)
	fmt.Printf("最后一个元素===> %v, 类型：%T \n", arr2[len(arr2)-1], arr2)

	var arr3 [5]int
	fmt.Println("arr3 ====>", arr3)

	// 让编译器自动推断长度
	var a6 = [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println("a6============>", a6)
	fmt.Println("a6============>", len(a6))

	//切片
	var a5 []string
	fmt.Println("a5============>", a5, a5 == nil)

	a5 = make([]string, 5)
	fmt.Println("a5============>", a5)

	fmt.Printf("%T \n", a5)
	a5 = append(a5, "1")
	a5 = append(a5, "1")
	fmt.Println("a5============>", a5, len(a5))

}
