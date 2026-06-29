package models

// 文章模型
type Post struct {
}

func (Post) TableName() string {
	return "post"
}
