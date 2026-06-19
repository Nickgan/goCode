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
	//fmt.Println("请输入你的年龄：")
	//var age int
	//fmt.Scan(&age)
	//
	//switch {
	//case age <= 0:
	//	fmt.Println("未出生")
	//	fallthrough
	//case age <= 18:
	//	fmt.Println("未成年")
	//	fallthrough
	//case age <= 35:
	//	fmt.Println("青年")
	//default:
	//	fmt.Println("中年")
	//}

	//mm1()
	mm2()
}

func mm1() {

	var age int
	fmt.Print("请输入你的年龄：")
	fmt.Scan(&age)
	switch {
	case age <= 18:
		fmt.Println("未成年..")
		fmt.Println("未成年122")
	case age <= 50:
		fmt.Println("青年")
		fmt.Println("青年222")
		fallthrough
	default:
		fmt.Println("老年")

	}
}

func mm2() {
	var week int
	fmt.Print("请输入星期数：")
	n, er := fmt.Scan(&week)

	if er != nil && n != 1 {
		fmt.Println("输入错误 err")
		return
	}
	switch week {
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	case 4:
		fmt.Println("星期四")
	case 5:
		fmt.Println("星期五")
	case 6:
		fmt.Println("星期六")
	case 7:
		fmt.Println("星期日")
	default:
		fmt.Println("输入错误")
	}
}
