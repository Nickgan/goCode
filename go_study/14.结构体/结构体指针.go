package main

import (
	"encoding/json"
	"fmt"
)

type UserInfo struct {
	Name string `json:"name"`
}

func (u *UserInfo) SetName(name string) {
	fmt.Printf("setName====>%p \n", u)
	u.Name = name
}

func main() {
	var u = UserInfo{
		Name: "张三",
	}
	u.SetName("李四")
	//fmt.Println(u.Name)

	bytedata, _ := json.Marshal(u)
	fmt.Println(string(bytedata))

	fmt.Printf("main=====>%p \n", &u)
}
