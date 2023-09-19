package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	PostID      int64  `json:"id" db:"post_id"`                                   //帖子ID
	AuthorID    int64  `json:"author_id" db:"author_id"`                          //作者ID
	CommunityID int64  `json:"community_id" db:"community_id" binding:"required"` //所属的社区
	Status      int32  `json:"status" db:"status"`                                //帖子的状态
	Title       string `json:"title" db:"title" binding:"required"`               //标题
	Content     string `json:"content" db:"content" binding:"required"`           //帖子内容
}

type ApiPost struct {
	CreatedAt   time.Time
	AuthorID    int64  `json:"author_id"`    //作者ID
	CommunityID int64  `json:"community_id"` //所属的社区
	Title       string `json:"title" `       //标题
	Content     string `json:"content"`      //帖子内容
}

// ApiPostDetail 帖子详情接口的结构体
type ApiPostDetail struct {
	AuthorName          string `json:"author_name"` //作者姓名
	*ApiPost            `json:"post"`
	*ApiCommunityDetail `json:"community"` // 嵌入社区信息
}
