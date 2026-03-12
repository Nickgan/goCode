package main

import "fmt"

func main() {

	var str string = "hello world 成都"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}

	// 如果字符串中含有中文，那么上面的遍历就不行了，必须用下面的range遍历
	for i, v := range str {
		fmt.Printf("%d:%c\n", i, v)
	}
}
