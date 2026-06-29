package models

import (
	"fmt"
	"goCode/go_gorm/models/dto"
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	ID         int64          `gorm:"primaryKey"`      // 主键
	Name       string         `gorm:"not null;unique"` //不能为空，唯一
	Age        int            `gorm:"not null"`
	Info       dto.UserInfo   `gorm:"column:userInfo;type:json;serializer:json"`
	CreateTime time.Time      `gorm:"autoCreateTime"` // 在创建记录时自动设置为当前时间
	DeletedAt  gorm.DeletedAt // 软删除
	Status     dto.UserStatus `gorm:"column:status;type:json;serializer:json"`
	Money      int            `gorm:"column:money;default:0"`
}

// BeforeCreate 创建(insert)的钩子函数
func (u *UserModel) BeforeCreate(tx *gorm.DB) error {
	fmt.Println("创建的钩子函数")

	fmt.Println("枫枫" + time.Now().Format("2006-01-02 15:04:05.000"))
	u.Name = "枫枫" + time.Now().Format("2006-01-02 15:04:05.000")
	return nil
}

func (u *UserModel) BeforeUpdate(tx *gorm.DB) (err error) {
	fmt.Println("更新的钩子函数")
	return nil
}

func (u *UserModel) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Println("删除钩子")
	return
}

func (u *UserModel) AfterFind(tx *gorm.DB) (err error) {
	fmt.Println("查询钩子")
	return
}
