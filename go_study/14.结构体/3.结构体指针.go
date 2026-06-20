package main

import (
	"fmt"
)

type UserInfo struct {
	Name string `json:"name"`
}

func (u *UserInfo) SetName(name string) {
	u.Name = name
}

func (u *UserInfo) GetName() string {
	return u.Name
}

func main() {
	var u = UserInfo{
		Name: "张三",
	}
	u.SetName("李四")
	fmt.Println(u.Name)
	fmt.Println(u.GetName())

	//bytedata, _ := json.Marshal(u)
	//fmt.Println(string(bytedata))
	//fmt.Printf("main=====>%p \n", &u)
}
