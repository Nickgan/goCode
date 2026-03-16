package main

import "fmt"

func main() {

	var map1 = map[int]string{
		1: "张三",
		2: "李四",
		3: "王五",
	}

	fmt.Println(map1)
	s, ok := map1[3]

	fmt.Printf("v===>%#v, 是否存在====>%v \n", s, ok)
}
