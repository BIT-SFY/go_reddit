package logic

import (
	"reddit/dao/mysql"
	"reddit/models"
	"reddit/pkg/snowflake"

	"go.uber.org/zap"
)

// CreatePost 创建帖子
func CreatePost(p *models.Post) (err error) {
	//1. 生成post id
	p.PostID = snowflake.GenID()
	//2. 保存到数据库,并返回错误
	return mysql.CreatePost(p)
}

// GetPostById 通过帖子id找到帖子具体内容
func GetPostById(pid int64) (data *models.ApiPostDetail, err error) {
	// 查询并组合我们接口想用的数据
	post, err := mysql.GetPostById(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostById(pid) failed",
			zap.Int64("pid", pid),
			zap.Error(err))
		return
	}

	// 根据作者id查询作者信息
	user, err := mysql.GetUerById(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUerById(post.AuthorID)",
			zap.Int64("author_id", post.AuthorID),
			zap.Error(err))
		return
	}
	// 根据社区id查询社区详细信息
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID)",
			zap.Int64("community_id", post.CommunityID),
			zap.Error(err))
		return
	}
	// 接口数据拼接
	data = &models.ApiPostDetail{
		AuthorName:         user.Username,
		ApiPost:            post,
		ApiCommunityDetail: community,
	}
	return
}

// GetPostList 获取帖子列表
func GetPostList(page, size int64) ([]*models.ApiPostDetail, error) {
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		return nil, err
	}
	data := make([]*models.ApiPostDetail, 0, len(posts))
	// 为所有帖子细节进行补充
	for _, post := range posts {
		// 根据作者id查询作者信息
		user, err := mysql.GetUerById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUerById(post.AuthorID)",
				zap.Int64("author_id", post.AuthorID),
				zap.Error(err))
			return nil, err
		}
		// 根据社区id查询社区详细信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID)",
				zap.Int64("community_id", post.CommunityID),
				zap.Error(err))
			return nil, err
		}
		// 接口数据拼接
		postDetail := &models.ApiPostDetail{
			AuthorName:         user.Username,
			ApiPost:            post,
			ApiCommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return data, err
}
