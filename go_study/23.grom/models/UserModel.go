package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	ID         int64     `gorm:"primaryKey"`      // 主键
	Name       string    `gorm:"not null;unique"` //不能为空，唯一
	Age        int       `gorm:"not null"`
	CreateTime time.Time `gorm:"autoCreateTime"` // 在创建记录时自动设置为当前时间
}

// BeforeCreate 创建(insert)的钩子函数
func (u *UserModel) BeforeCreate(tx *gorm.DB) error {
	fmt.Println("创建的钩子函数")

	fmt.Println("枫枫" + time.Now().Format("2006-01-02 15:04:05.000"))
	u.Name = "枫枫" + time.Now().Format("2006-01-02 15:04:05.000")
	time.Sleep(2 * time.Second)
	return nil
}

func (u *UserModel) BeforeUpdate(tx *gorm.DB) (err error) {
	fmt.Println("更新的钩子函数")
	return nil
}
