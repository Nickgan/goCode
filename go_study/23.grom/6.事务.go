package main

import (
	"fmt"

	"goCode/go_study/23.grom/global"
	"goCode/go_study/23.grom/models"
	"gorm.io/gorm"
)

func main() {

	tx1()
	//tx2()
}

// 写法一：自动事务(这种写法是写到一个函数里面，不用在里面去提交事务和回滚事务)
func tx1() {
	// 张三向李四转账100元
	global.ConnectDB()
	//global.DB.AutoMigrate(&models.UserModel{})

	var zs models.UserModel
	var lisi models.UserModel
	global.DB.Where("name = ? ", "张三").Find(&zs)
	global.DB.Where("name = ? ", "李四").Find(&lisi)

	global.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&models.UserModel{}).Where("id = ?", zs.ID).Update("money", zs.Money-100).Error
		if err != nil {
			fmt.Println("更新张三账户出错了", err)
		}

		err = tx.Model(&models.UserModel{}).Where("id = ?", lisi.ID).Update("money", lisi.Money+100).Error
		if err != nil {
			fmt.Println("更新李四账户出错了")
		}

		// 提交事务
		return nil
	})
}

// 写法二：手动事务（需要手动提交和回滚事务）
func tx2() {
	// 张三向李四转账100元
	global.ConnectDB()
	//global.DB.AutoMigrate(&models.UserModel{})

	var zs models.UserModel
	var lisi models.UserModel
	global.DB.Where("name = ? ", "张三").Find(&zs)
	global.DB.Where("name = ? ", "李四").Find(&lisi)

	tx := global.DB.Begin()

	err := tx.Model(&models.UserModel{}).Where("id = ?", zs.ID).Update("money", zs.Money-100).Error
	if err != nil {
		fmt.Println("更新张三账户出错了", err)
		tx.Rollback()
	}

	err = tx.Model(&models.UserModel{}).Where("id = ?", lisi.ID).Update("money", lisi.Money+100).Error
	if err != nil {
		fmt.Println("更新李四账户出错了")
		tx.Rollback()
	}

	tx.Commit()
}
