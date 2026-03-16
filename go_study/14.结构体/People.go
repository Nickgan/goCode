package main

import "fmt"

type Person struct {
	Time string
}

func (p Person) PrintInfo() {
	fmt.Printf("Time:%s\n", p.Time)
}

// Student 定义结构体
type Student struct {
	Person
	Name string
	Age  int
}

// PrintInfo 给结构体添加一个方法
func (s Student) PrintInfo() {
	fmt.Printf("Name:%s, Age:%d\n", s.Name, s.Age)
}

func main() {
	s := Student{Name: "张三", Age: 18, Person: Person{Time: "2020-01-01"}}

	s.Name = "李四"
	s.PrintInfo()
	s.Person.PrintInfo()
}
