package models

// 评论模型
type Comment struct {
}

func (Comment) TableName() string {
	return "comment"
}
