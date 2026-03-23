package main

import "fmt"

type Person interface {
	say()
}

type Man struct {
	Name string
}

func (m Man) say() {
	fmt.Printf("%v 说，你好Man \n", m.Name)
}

type Woman struct {
	Name string
}

func (w Woman) say() {
	fmt.Printf("%v 说，你好 Woman \n", w.Name)
}

func main() {
	m := Man{
		Name: "张三",
	}

	w := Woman{
		Name: "小丽",
	}

	var p Person = m
	var p1 Person = w

	p.say()
	p1.say()

	// 断言类型
	switch p1.(type) {
	case Man:
		fmt.Println("Man类型")
	case Woman:
		fmt.Println("Woman类型")
	}

	// 断言某个类型
	c, ok := p.(Man) // 两个参数    断言之后的类型   是否是对应类型
	fmt.Println(c, ok)

	c1, ok1 := p1.(Woman) // 两个参数    断言之后的类型   是否是对应类型
	fmt.Println(c1, ok1)
}
