package main

import "fmt"

// Code 自定义一个基本类型的主要目的就是为了方便绑定方法，比如这里的GetMsg
type Code int

var (
	OK    Code = 200
	Error Code = 500
)

func (c Code) GetMsg() {
	switch c {
	case OK:
		fmt.Println("成功")
	case Error:
		fmt.Println("失败")
	}
}

func main() {
	var code int

	fmt.Print("请输入状态码：")
	fmt.Scan(&code)

	if code == 200 {
		OK.GetMsg()
	} else if code == 500 {
		Error.GetMsg()
	} else {
		fmt.Println("请输入正确的状态码")
	}
}
