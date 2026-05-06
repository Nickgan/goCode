package main

import "fmt"

func main1() {

	//<=0     未出生
	//1-18    未成年
	//18-35   青年
	//>=35    中年

	var age int
	fmt.Print("请输入你的年龄:")

	_, error := fmt.Scan(&age)

	if error != nil {
		fmt.Println("输入错误")
		return
	}

	switch {
	case age <= 0:
		fmt.Println("未出生")
	case age <= 18:
		fmt.Println("未成年")
	case age <= 35:
		fmt.Println("青年")
	default:
		fmt.Println("中年")

	}

}

func main() {
	fmt.Println("请输入你的年龄：")
	var age int
	fmt.Scan(&age)

	switch {
	case age <= 0:
		fmt.Println("未出生")
		fallthrough
	case age <= 18:
		fmt.Println("未成年")
		fallthrough
	case age <= 35:
		fmt.Println("青年")
	default:
		fmt.Println("中年")
	}
}
