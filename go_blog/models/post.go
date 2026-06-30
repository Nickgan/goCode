package models

import (
	"time"

	"gorm.io/gorm"
)

// Post 文章模型
type Post struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title" gorm:"not null"`
	Content   string         `json:"content" gorm:"type:text;not null"`
	UserID    int            `json:"user_id" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (Post) TableName() string {
	return "post"
}
