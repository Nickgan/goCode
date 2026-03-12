package main

import (
	"fmt"
)

func main() {

	f3(6)
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
