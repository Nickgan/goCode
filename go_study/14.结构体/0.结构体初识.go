package main

import "fmt"

type User struct {
	Name string
}

func (u User) printUserInfo() {
	fmt.Println("printUserInfo============>", u)
}

func main() {
	var u = User{
		Name: "张三",
	}

	u.printUserInfo()
	u.Name = "aaa"
	u.printUserInfo()
}
