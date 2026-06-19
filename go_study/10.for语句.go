package main

import "fmt"

func main() {

	//for1()
	f99()

	// 每隔一秒打印时间
	//for {
	//	fmt.Println(time.Now().Format("2006_01-02 15:04:05"))
	//	time.Sleep(time.Second)
	//}
}

func for1() {

	//var sum int
	//for i := 1; i <= 100; i++ {
	//	sum += i
	//}
	//fmt.Println(sum)

	// while 模式
	//i := 0
	//sum := 0
	//for i < 100 {
	//	i++
	//	sum += i
	//}
	//fmt.Println(sum)

	// do while模式
	//i := 0
	//sum := 0
	//for {
	//	i++
	//	sum += i
	//	if i >= 100 {
	//		break
	//	}
	//}
	//fmt.Println(sum)

	// 遍历切片
	var list = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range list {
		fmt.Println(list[i])
	}
	fmt.Println("===========")
	for i, s := range list {
		fmt.Println(i, s)
	}
	fmt.Println("===========")

	// 遍历map
	var nameMap = map[int]string{
		1: "张三",
		2: "李四",
		3: "王五",
	}
	for i, s := range nameMap {
		fmt.Println(i, s)
	}
	fmt.Println("===========")
	for i := range nameMap {
		fmt.Println(i, nameMap[i])
	}
}

// 九九乘法表
func f99() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			//if j > i {
			//
			//	break
			//}
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
		}
		fmt.Println()
	}
}
