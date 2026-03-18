package main

import (
	"encoding/json"
	"fmt"
)

/*
*
  - 结构体Tag

*这个不写json标签转换为json的话，字段名就是属性的名字
小写的属性也不会转换
*/
type Student1 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  bool   `json:"-"` //中划线标识，不转换

	//Name string
	//Age  int
}

func main() {

	var student = Student1{Name: "张三", Age: 18, Sex: true}
	byteData, _ := json.Marshal(student)
	fmt.Println(string(byteData))
}
