package main

import "fmt"

type PersonNew struct {
	Name  string
	Age   int
	Sex   string
	Fight int
}

func (p *PersonNew) SetAge(age int) {
	p.Age = age
	fmt.Println("PersonNew SetAge =========>", age)
}

type SuperMan struct {
	Strength int
	Speed    int
	ps       PersonNew
}

func (p *SuperMan) SetAge(age int) {
	p.ps.Age = age
	fmt.Println("SuperMan SetAge =========>", age)
}

func (s *SuperMan) print() {
	fmt.Println("SuperMan print====>", s)
}

func main() {
	p1 := PersonNew{Name: "张三", Age: 18, Sex: "男", Fight: 90}
	(&p1).SetAge(20)

	s1 := SuperMan{Strength: 100, Speed: 80, ps: p1}
	//s1.print()

	s1.SetAge(25)

	s1.ps.SetAge(30)

}
