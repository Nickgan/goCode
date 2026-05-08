package main

import "fmt"

// Animal 定义一个Animal接口，他有唱、跳、rap三个方法
type Animal interface {
	sing()
	jump()
	rap()
}

// Chicken 定义一个Chicken结构体,需要实现这些方法
type Chicken struct {
	Name string
}

func (c Chicken) sing() {
	fmt.Printf("%v 正在唱歌\n", c.Name)
}

func (c Chicken) jump() {
	fmt.Printf("%v 正在跳\n", c.Name)
}

func (c Chicken) rap() {
	fmt.Printf("%v 正在rap\n", c.Name)
}

type Cat struct {
	Name string
}

func (c Cat) sing() {
	fmt.Printf("%v 正在唱歌\n", c.Name)
}
func (c Cat) jump() {
	fmt.Printf("%v 正在跳\n", c.Name)
}
func (c Cat) rap() {
	fmt.Printf("%v 正在rap\n", c.Name)
}

// 多态方法
func sing(a Animal) {
	a.sing()
}

// 全部实现完之后，chicken就不再是一只普通的鸡了

func main() {
	// 创建一个Chicken结构体变量
	c := Chicken{Name: "小鸡"}
	cat := Cat{Name: "小猫"}

	sing(c)
	sing(cat)

	fmt.Println()
}
