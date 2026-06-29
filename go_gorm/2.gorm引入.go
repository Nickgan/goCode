package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID   uint
	Name string
	Age  int
}

func (User) TableName() string {
	return "users"
}

func main() {
	// 连接数据库
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败")
	} else {
		fmt.Println("数据库连接成功")
	}

	var userList []User
	db.Find(&userList)

	for _, user := range userList {
		fmt.Println(user.ID, user.Name, user.Age)
	}
}
