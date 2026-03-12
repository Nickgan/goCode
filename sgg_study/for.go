package main

import (
	"fmt"
)

func main() {

	f4(10)

	//f2()
}

/*
*
打印空心金字塔
*/
func f4(n int) {

	var totalLevel int = n

	// 1 表示层数
	for i := 1; i <= totalLevel; i++ {

		// 打印*前面的空格
		for j := 1; j <= totalLevel-i; j++ {
			fmt.Print(" ")
		}

		// 2 层数i，打印*
		for k := 1; k <= 2*i-1; k++ {
			if k == 1 || k == 2*i-1 || i == totalLevel {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}

}

func f3(n int) {
	for i := 0; i <= n; i++ {

		var m int = i
		var k = n - i

		fmt.Printf("%d + %d = %d\n", m, k, n)
	}
}

func f2() {

	var sum int
	for i := 1; i <= 100; i++ {
		if i%9 == 0 {
			fmt.Println(i)
			sum += i
		}
	}

	fmt.Println("sum=", sum)

}

func f1() {

	var str string = "hello world成都"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}

	// 如果字符串中含有中文，那么上面的遍历就不行了，必须用下面的range遍历,或者是用下面第二种：“切片遍历”
	for i, v := range str {
		fmt.Printf("%d:%c\n", i, v)
	}

	//转换切片类型，“切片遍历”
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c\n", str2[i])
	}
}
