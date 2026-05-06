package main

import (
	"fmt"
	"time"
)

func main() {

	//var arr [3]int = [3]int{1, 2, 3}
	//fmt.Println(arr)
	//
	//var arr2 = [3]int{1, 2, 5}
	//
	//arr2[2] = 10
	//fmt.Println(arr2)
	//
	//fmt.Println(len(arr2))
	//
	////声明一个切片slice
	//var s1 []string
	//s1 = make([]string, 0)
	//s1 = make([]string, 0, 10)
	//
	//fmt.Println(s1 == nil)

	//f1()

	var sum int
	for i := 1; i <= 100; i++ {
		sum += i
		fmt.Println("i==========>", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("一到100的和为：", sum)
}

func f1() {
	var week = [7]string{"星期一", "星期二", "星期三", "星期四", "星期五", "星期六", "星期日"}

	var weekNumber int
	fmt.Println("请输入数字：")
	fmt.Scan(&weekNumber)

	if weekNumber < 1 || weekNumber > 7 {
		fmt.Println("输入的数字有误")
		f1()
		return
	}
	fmt.Println(week[weekNumber-1])
}
