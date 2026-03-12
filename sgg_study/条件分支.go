package main

import "fmt"

func main() {

	//m1()

	m3()
}

func m3() {
	var key byte
	fmt.Print("请输入一个字符：")
	fmt.Scanf("%c", &key)
	switch key {
	case 'a':
		fmt.Println("aaaa111")
		fmt.Println("aaaa222")
	case 'b':
		fmt.Println("bbbb")
	default:
		fmt.Println("输入错误")
	}

}

func m1() {
	fmt.Print("请输入一个分数：")

	var score int
	fmt.Scan(&score)

	if score >= 100 {
		fmt.Println("奖励一辆宝马")
	} else if score > 80 && score <= 99 {
		fmt.Println("奖励一台iphone")
	} else if score > 60 && score <= 80 {
		fmt.Println("奖励一台ipad")
	} else {
		fmt.Println("什么奖励都没有")
	}
}

func m2() {

	var n int = 10

	if n > 9 {
		fmt.Println("ok1")
	} else if n > 6 {
		fmt.Println("ok2")
	} else if n > 3 {
		fmt.Println("ok3")
	} else {
		fmt.Println("ok4")
	}
}
