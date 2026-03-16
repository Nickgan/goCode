package main

import "fmt"

func main() {

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

	m2(age)
}

func m2(age int) {

	switch {
	case age <= 0:
		fmt.Println("未出生")
		fallthrough
	case age <= 18:
		fmt.Println("未成年")
		fallthrough
	case age <= 35:
		fmt.Println("青年")
	case age >= 35:
		fmt.Println("中年")
	}
}

func m1(age int) {
	if age <= 0 {
		fmt.Println("未出生")
	} else if age > 0 && age < 18 {
		fmt.Println("未成年")
	} else if age >= 18 && age < 35 {
		fmt.Println("青年")
	} else if age >= 35 {
		fmt.Println("中年")
	}
}
