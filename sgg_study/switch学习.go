package main

import "fmt"

func main() {

	//m1()

	m2()
}

func m11() {

	var a int = 5566

	switch a {
	case 1, 2, 10:
		fmt.Println("a is 1 or 2 or 10")
	case 55:
		fmt.Println("a is 55")
	default:
		fmt.Println("default================")

	}
}

func m21() {

	fmt.Println("请输入一个字符，a,b,c,d,e,f,g")
	var a byte
	fmt.Scanf("%c", &a)
	switch a {
	case 'a', 'b', 'c', 'd', 'e', 'f', 'g':
		fmt.Println("输入的字符是：", a)
	default:
		fmt.Println("输入的字符不是：a,b,c,d,e,f,g")
	}

}
