package main

import "fmt"

func main() {
	fmt.Printf("%v\n", "你好")          // 可以作为任何值的占位符输出
	fmt.Printf("%v %T\n", "枫枫", "枫枫") // 打印类型
	fmt.Printf("%d\n", 3)             // 整数
	fmt.Printf("%.2f\n", 1.2511)      // 小数
	fmt.Printf("%s\n", "哈哈哈")         // 字符串
	fmt.Printf("%#v\n", "aaa")

}
