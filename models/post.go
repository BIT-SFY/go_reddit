package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	PostID      int64  `json:"id" db:"post_id"`                //帖子ID
	Title       string `json:"title" db:"title"`               //标题
	Content     string `json:"content" db:"content"`           //帖子内容
	AuthorID    int64  `json:"author_id" db:"author_id"`       //作者ID
	CommunityID int64  `json:"community_id" db:"community_id"` //所属的社区
	Status      int64  `json:"status" db:"status"`             //帖子的状态
}
