package main

import (
	"encoding/json"
	"fmt"

	"goCode/go_study/23.grom/global"
	"goCode/go_study/23.grom/models"
)

func main() {

	//user := models.UserModel{
	//	Name: "王五111",
	//	Age:  18,
	//	Info: dto.UserInfo{
	//		Like:    []string{"篮球", "游戏"},
	//		Address: "北京",
	//	},
	//	CreateTime: time.Now(),
	//	Status:     dto.UserStatusNormal,
	//}
	//
	global.ConnectDB()
	//global.DB.AutoMigrate(&models.UserModel{})
	//global.DB.Create(&user)
	//fmt.Println("用户创建成功")

	// 查询
	var userDb models.UserModel
	global.DB.Debug().Where("id = ?", 65).Find(&userDb)
	fmt.Println("查询结果:", userDb)
	fmt.Println("用户状态:", userDb.Status)

	b, _ := json.Marshal(userDb)
	fmt.Println(string(b))
}
