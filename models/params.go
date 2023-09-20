package models

const (
	OrderTime  = "time"
	OrderScore = "score"
)

//定义请求的参数的结构体

// ParamSignUp 注册参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 登录请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamVoteData 投票数据
type ParamVoteData struct {
	// UserID    从请求种获取当前的用户
	PostID    int64 `json:"post_id" binding:"required"`       //帖子ID
	Direction int8  `json:"direction" binding:"oneof=-1 0 1"` //赞成(1)反对(-1)取消投票(0)
}

// ParamPostList 获取帖子列表query string参数
type ParamPostList struct {
	Page  int64  `json:"page" form:"page"`
	Size  int64  `json:"size" form:"size"`
	Order string `json:"order" form:"order"`
}
