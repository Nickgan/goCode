package main

import "fmt"

func main() {
	fmt.Printf("%v\n", "你好")          // 可以作为任何值的占位符输出
	fmt.Printf("%v %T\n", "枫枫", "枫枫") // 打印类型
	fmt.Printf("%d\n", 3)             // 整数
	fmt.Printf("%.2f\n", 1.25)        // 小数
	fmt.Printf("%s\n", "哈哈哈")         // 字符串
	fmt.Printf("%#v\n", "哇。。。")       // 用go的语法格式输出，很适合打印空字符串

	fmt.Printf("%v,%v,%v", "aaa", 12.5, 123)

	fmt.Println("=====================")

	m11111()
}

func m11111() {

	fmt.Printf("%v \n", "任意类型打印")
	fmt.Printf("%d \n", 23)
	fmt.Printf("%.3f \n", 23.5)
	fmt.Printf("%.2f \n", 23.1)
	fmt.Printf("%#v \n", "用go的语法输出")
	fmt.Printf("%v  %T \n", "任意类型打印", "任意类型打印")

	//付给变量
	var amount = fmt.Sprintf("%.2f", 23.666)
	fmt.Println(amount)

	// 输入
	var age11 int
	var e1 error
	fmt.Print("请输入你的年龄：")
	fmt.Scan(&age11)
	fmt.Println("你的年龄是：", age11)
	if e1 != nil {
		fmt.Println("输入错误")
	}

}
