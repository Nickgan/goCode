package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Time string `json:"time_"`
	Name string
	age  int `json:"age"`
}

func (p Person) info() {
	fmt.Println("Person.info====>", p)
}

func (p Person) setName(name string) {
	p.Name = name
	fmt.Println(&p)

}

func (p *Person) setName2(name string) {
	p.Name = name

}

type Student struct {
	Person
	Name string
}

func main() {
	p := Person{Time: "2020-01-01", Name: "张三", age: 18}
	fmt.Println(p.Name)
	fmt.Println(&p)

	p.setName("张三1111")
	fmt.Println(p.Name)

	p.setName2("张三2222")
	fmt.Println(p.Name)

	byteData, _ := json.Marshal(p)
	fmt.Println(string(byteData))
	fmt.Println(p.age)
	fmt.Println(p)
}
