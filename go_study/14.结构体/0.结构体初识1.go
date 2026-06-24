package main

import "fmt"

type Person1 interface {
	say()
}

type User1 struct {
	Name string
}

//func (u User1) say() {
//	fmt.Println("say============>", u)
//}

func (u *User1) say() {
	fmt.Println("say1============>", u)
}

func work(p Person1) {
	p.say()
}
func main() {
	var u = User1{
		Name: "张三",
	}

	(&u).say()
	u.say()

	//work(&u)
	//work(u)
}
