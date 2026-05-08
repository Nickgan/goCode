package main

import (
	"encoding/json"
	"fmt"
)

/*
*
  - 结构体Tag

1: 这个不写json标签转换为json的话，字段名就是属性的名字
2: 小写的属性也不会转换
3: 如果再转json的时候，我不希望某个字段被转出来，我可以写一个 -
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
