package main

import "encoding/json"

/*
*
  - 结构体Tag

*这个不写json标签转换为json的话，字段名就是属性的名字
小写的属性也不会转换
*/
type Student1 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`

	//Name string
	//Age  int
}

func main() {

	var student = Student1{Name: "张三", Age: 18}
	byteData, _ := json.Marshal(student)
	println(string(byteData))
}
