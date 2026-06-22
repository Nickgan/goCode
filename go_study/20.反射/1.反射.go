package main

import (
	"fmt"
	"reflect"
)

// 1：类型判断
func reFType(o any) {
	oType := reflect.TypeOf(o)
	switch oType.Kind() {
	case reflect.Int:
		fmt.Println("int")
	case reflect.String:
		fmt.Println("string")
	case reflect.Bool:
		fmt.Println("bool")
	case reflect.Map:
		fmt.Println("map")
	case reflect.Slice:
		fmt.Println("slice")
	case reflect.Struct:
		fmt.Println("struct")
	default:
		fmt.Println("other")
	}
}

// 通过反射获取值
func reFValue(o any) {
	v := reflect.ValueOf(o)
	fmt.Println(v, v.Type())
	switch v.Kind() {
	case reflect.Int:
		fmt.Println(v.Int())
	case reflect.String:
		fmt.Println(v.String())
	case reflect.Bool:
		fmt.Println(v.Bool())
	case reflect.Map:
		fmt.Println(v.MapIndex(reflect.ValueOf("a")))
	case reflect.Slice:
		fmt.Println(v)
	case reflect.Struct:
		fmt.Println(v.Interface())
	default:
		fmt.Println("other")
	}
}

// 3: 通过反射修改值
func reFChange(v1 any, v2 any) {
	t1 := reflect.TypeOf(v1)
	t2 := reflect.TypeOf(v2)

	// 类型一致性判断(这里v1是指针类型，要调用elem（）方法解引用后判断)
	if t1.Elem().Kind() != t2.Kind() {
		fmt.Println("类型不一致")
		return
	}

	vo := reflect.ValueOf(v1)
	elem := vo.Elem() //解引用
	switch elem.Kind() {
	case reflect.String:
		elem.SetString(reflect.ValueOf(v2).String())
		fmt.Println(v1)
	case reflect.Int:
		elem.SetInt(reflect.ValueOf(v2).Int())
		fmt.Println(v1)
	}
}

func main() {

	// 1：类型判断
	var a = 1
	reFType(a)
	var b = "hello"
	reFType(b)
	var c = true
	reFType(c)
	var d = map[string]int{"a": 1}
	reFType(d)
	var e = []int{1, 2, 3}
	reFType(e)
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var f = User{Name: "张三", Age: 18}
	reFType(f)

	// 2：通过反射获取值
	reFValue(f)

	// 3: 通过反射修改值
	reFChange(&b, "hello world11111111111")
	fmt.Println("b=============>", b)
}
