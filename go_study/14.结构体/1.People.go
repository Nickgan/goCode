package main

import "fmt"

type Person struct {
	Time string
	Name string
}

// Info 给结构体添加一个方法
func (p Person) Info() {
	fmt.Printf("Time:%s\n", p.Time)
}

// Student 定义结构体
type Student struct {
	Person // 这里组合就是java的继承
	Name   string
	Age    int
}

// PrintInfo 给结构体添加一个方法
func (s Student) PrintInfo() {
	fmt.Printf("Name:%s, Age:%d\n", s.Name, s.Age)
}

// SetName 给结构体添加一个方法（验证值传递）
func (s Student) SetName(name string) {
	fmt.Printf("SetName， s内存地址：%p \n", &s)
	s.Name = name
}

// SetName2 给结构体添加一个方法（验证指针传递）
func (s *Student) SetName2(name string) {
	fmt.Printf("SetName2， s内存地址：%p \n", &s)
	s.Name = name
}

func main() {
	s := Student{Name: "张三", Age: 18, Person: Person{Time: "2020-01-01", Name: "张三P"}}
	fmt.Printf("main， s内存地址：%p \n", &s)
	s.SetName("张三。。。。")
	fmt.Println(s.Name)

	s.SetName2("张三。。。。")
	fmt.Println(s.Name)

	//s.Name = "李四"
	//s.PrintInfo()
	//
	////调用父结构体的方法
	//s.Info()
	//
	//// 调用
	//s.Person.Info()
	//
	////访问父结构体的属性
	//fmt.Println(s.Time)
	//
	//fmt.Println(s.Person.Time)
	//
	//fmt.Println(s.Person.Name)
	//fmt.Println(s.Name)
}
