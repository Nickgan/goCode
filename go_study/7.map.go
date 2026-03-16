package main

import "fmt"

func main() {

	// 声明 1
	var map1 = map[int]string{
		1: "张三",
		2: "李四",
		3: "王五",
	}

	// 声明 2
	var map2 = make(map[string]string)
	fmt.Println("map2=====>", map2)

	// 取值
	fmt.Println(map1)
	s, ok := map1[3]
	fmt.Printf("v===>%#v, 是否存在====>%v \n", s, ok)

	//设置值
	map1[4] = "赵六"
	fmt.Println(map1)

	//删掉值
	delete(map1, 4)
	fmt.Println(map1)

	fmt.Println("----------------------map的遍历------------------")
	for k, v := range map1 {
		fmt.Printf("k==>%v, v==>%v \n", k, v)
	}

	// map是引用传递，数组是值传递，切片是引用传递
	var m1 = map[int]string{
		1: "张三",
		2: "李四",
		3: "王五",
	}

	fmt.Println(m1[1])
	mapyy(m1)
	fmt.Println(m1[1])
}

func mapyy(m map[int]string) {
	m[1] = "hello world"
}
