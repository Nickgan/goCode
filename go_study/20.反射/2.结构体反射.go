package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"username"`
	Age  int    `json:"age"`
}

func main() {

	// 读取结构体的属性
	var u = User{
		Name: "张三",
		Age:  18,
	}

	v := reflect.ValueOf(u)
	t := reflect.TypeOf(u)

	fmt.Println("结构体类型名：", t.Name())
	fmt.Println("kind:", t.Kind())
	fmt.Println("值：", v)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("字段名：%s, tag: %s, 类型：%s\n", field.Name, field.Tag.Get("json"), field.Type)
	}

	// 修改结构体的属性
	var u1 = User{}
	userName := reflect.ValueOf(&u1).Elem().FieldByName("username")
	if userName.CanSet() {
		userName.SetString("甘波")
	}

	ave := reflect.ValueOf(&u1).Elem().FieldByName("age")
	if ave.CanSet() {
		ave.SetInt(18)
	}

	fmt.Println(u1)

	jsonStr, _ := json.Marshal(u1)
	fmt.Println(string(jsonStr))
}
