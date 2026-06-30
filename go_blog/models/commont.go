package models

import (
	"time"

	"gorm.io/gorm"
)

// 评论模型
type Comment struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	Content   string         `json:"content" gorm:"type:text;not null"`
	UserID    int            `json:"user_id" gorm:"not null"`
	PostID    int            `json:"post_id" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (Comment) TableName() string {
	return "comment"
}
