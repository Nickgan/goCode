package models

// 用户模型
type User struct {
}

func (User) TableName() string {
	return "user"
}
