package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) printAnimalName() {
	fmt.Println("AnimalName=============>", a.Name)
}

type Cat struct {
	Animal //继承会继承父类的属性和方法
	Age    int
}

func main() {
	var cat = Cat{
		Age: 18,
		Animal: Animal{
			Name: "小猫",
		},
	}

	cat.printAnimalName()
}
