package main

import "fmt"

func main() {
	//var a int = -10
	//var b int = -3
	//fmt.Println(a % b)

	//for i := 5; i < 20; i++ {
	//	fmt.Println(i)
	//}

	var days int = 97
	var week int = days / 7
	var day int = 97 % 7
	fmt.Println(week, day)

	var huashi = 134.2
	var sheshi = 5.0 / 9 * (huashi - 100)
	fmt.Println(sheshi)

	fmt.Printf("%v", sheshi)
	fmt.Println("%.2f", sheshi)

	fmt.Println("--------------关系运算符-------------------")

	var n1 int = 9
	var n2 int = 10
	fmt.Println(n1 == n2)
	fmt.Println(n1 != n2)
	fmt.Println(n1 > n2)
	fmt.Println(n1 < n2)
	fmt.Println(n1 >= n2)
	fmt.Println(n1 <= n2)
	flag := n1 > n2
	fmt.Println("flag====>", flag)
	fmt.Println("--------------逻辑运算符-------------------")
}
