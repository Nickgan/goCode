package main

import (
	"fmt"
	"goCode/go_gorm/global"
	"goCode/go_gorm/models"
	"time"

	"gorm.io/gorm"
)

func main() {
	// 连接数据库
	global.ConnectDB()

	// 设置自动创建表
	global.DB.AutoMigrate(&models.UserModel{})

	fmt.Println("生成表结构成功。")

	// 插入操作
	insertData(global.DB)

	// 查询操作
	selectData(global.DB)

	// 更新操作
	//updateData(global.DB)

	//删除操作
	//deleteData(global.DB)

	// 原生sql
	//rawSql(global.DB)
}

// 更新操作
func updateData(db *gorm.DB) {

	fmt.Println("1.============>save方法保存或者更新")
	var user = models.UserModel{
		ID:         61,
		Name:       "甘波111",
		Age:        181111,
		CreateTime: time.Now(),
	}

	err := db.Debug().Save(&user).Error
	if err != nil {
		fmt.Printf("操作失败，失败信息：%s", err.Error())
		return
	}
	fmt.Println("更新成功。=============>", user)

	fmt.Println("2.============>update方法更新")
	db.Model(&models.UserModel{}).Where("id", 61).Update("name", "甘波666")
	fmt.Println("Update更新成功。=============>", user)

	// 用法跟上面的uapdate一样，区别就是不会触发钩子函数
	db.Model(&models.UserModel{}).Where("id", 61).UpdateColumn("name", "甘波888")
	fmt.Println("UpdateColumn更新成功。=============>", user)

}

// 删除操作
func deleteData(db *gorm.DB) {
	fmt.Println("1.============>delete方法删除")
	db.Delete(&models.UserModel{}, 60)

	fmt.Println("2.============>批量删除")
	db.Delete(&models.UserModel{}, []int{61, 62, 63})

	// 查询删除记录
	var user models.UserModel
	db.Debug().Where("id = ?", 60).Find(&user)
	fmt.Println("查询删除记录。=============>", user)

	// Unscoped 忽略软删除记录查询
	db.Debug().Where("id = ?", 60).Unscoped().Find(&user)
	fmt.Println("忽略软删除记录查询。=============>", user)

	// 真正的物理删除
	db.Unscoped().Delete(&models.UserModel{}, 60)

}

// 查询操作
func selectData(db *gorm.DB) {
	var err error

	//fmt.Println("====查询全部=======")
	//var users []models.UserModel
	//err := db.Find(&users).Error
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//for _, user := range users {
	//	fmt.Println(user)
	//}

	fmt.Println("====查询单个(果查不到则会报错)=======")
	var user models.UserModel
	//err = db.Debug().Take(&user).Error
	err = db.Debug().Take(&user, "id=?", 100).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)

	fmt.Println("====条件查询=======")
	var users1 []models.UserModel
	err = db.Debug().Where("id = ?", 10).Find(&users1).Error
	if err != nil {
		fmt.Println(err)
	}
	for _, user := range users1 {
		fmt.Println(user)
	}

	fmt.Println("====条件查询2=======")
	var user2 []models.UserModel
	db.Debug().Find(&user2)
	fmt.Println("=============>", user2)
}

type userListDTO struct {
	ID       int
	UserName string `gorm:"column:name"`
}

// 原生 sql
func rawSql(db *gorm.DB) {
	var users []models.UserModel
	db.Debug().Raw("select * from user_models where deleted_at is null").Scan(&users)
	for i, v := range users {
		fmt.Println(i, v)
	}

	fmt.Println("查询部分字段============")
	var users1 []userListDTO
	db.Debug().Raw("select id,name from user_models where deleted_at is null").Scan(&users1)
	for i, v := range users1 {
		fmt.Println(i, v.UserName)
	}
}

func insertData(db *gorm.DB) {
	var users = []models.UserModel{
		{Name: "张三", Age: 18},
		//{Name: "李四", Age: 19},
		//{Name: "王五", Age: 20},
	}

	// 插入数据
	//err := db.Debug().Create(&users).Error
	err := db.Debug().Create(&users).Error //打印sql
	fmt.Println(users, err)

	// 创建成功之后，数据会回填
	fmt.Println("插入数据成功。")

}
