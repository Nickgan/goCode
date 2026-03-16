package main

import (
	"fmt"
	"sort"
	"strconv"
)

/*
切片对应java里面的List
*/
func main() {

	var list []int

	fmt.Println(list == nil)
	fmt.Printf("%T \n", list)

	list = append(list, 1)
	list = append(list, 2)

	fmt.Println(list)

	var nameList = make([]string, 10)
	fmt.Println("容量:", cap(nameList))
	for i := 0; i < 100; i++ {
		nameList = append(nameList, strconv.Itoa(i))
	}

	fmt.Println(len(nameList))
	fmt.Printf("长度====>%v,  内容:  %v \n", len(nameList), nameList)

	fmt.Println("容量:", cap(nameList))

	var arr []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 打印arr的类型
	fmt.Printf("%T \n", arr)

	s1 := arr[1:3]
	s1 = append(s1, 11)
	fmt.Printf("%T \n", s1)
	fmt.Println(s1)

	s2 := []int{1, 2, 3}
	fmt.Println(s2)

	var arr3 []string
	fmt.Println("arr3============>", arr3 == nil)

	//切片排序
	var arr4 = []int{1, 5, 3, 2, 4}
	fmt.Println("排序前:", arr4)
	sort.Ints(arr4)
	fmt.Println("排序后:", arr4)
	fmt.Println("从大到小排序")
	sort.Sort(sort.Reverse(sort.IntSlice(arr4)))
	fmt.Println(arr4)

	fmt.Println("==================数组切割成 子切片")
	// 数组切割成 子切片
	var arr2 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := arr2[2:3]
	fmt.Println(s)
}
