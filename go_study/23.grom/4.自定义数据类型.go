package main

import (
	"fmt"
	"time"

	"goCode/go_study/23.grom/global"
	"goCode/go_study/23.grom/models"
	"goCode/go_study/23.grom/models/dto"
)

func main() {

	user := models.UserModel{
		Name: "张三",
		Age:  18,
		Info: dto.UserInfo{
			Like:    []string{"篮球", "游戏"},
			Address: "北京",
		},
		CreateTime: time.Now(),
	}

	global.ConnectDB()
	global.DB.Create(&user)
	fmt.Println("用户创建成功")

	// 查询
	var userDb models.UserModel
	global.DB.Where("id = ?", 62).Find(&userDb)
	fmt.Println("查询结果:", userDb.Info.Like)
}
