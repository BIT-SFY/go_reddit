package logic

import (
	"reddit/dao/mysql"
	"reddit/dao/redis"
	"reddit/models"
	"reddit/pkg/snowflake"

	"go.uber.org/zap"
)

// CreatePost 创建帖子
func CreatePost(p *models.Post) (err error) {
	//1. 生成post id
	p.PostID = snowflake.GenID()
	//2. 保存到数据库,并返回错误
	err = mysql.CreatePost(p)
	if err != nil {
		return err
	}
	err = redis.CreatePost(p.PostID, p.CommunityID)
	if err != nil {
		return err
	}
	return
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

// GetPostList 获取帖子列表 升级版
func GetPostList(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	// 2.去redis查询id列表
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder(p) return 0 data")
		return
	}
	// 3.根据id去数据库查询帖子详情
	// 返回的数据还要按照给定的id的顺序返回
	posts, err := mysql.GetPostListByIds(ids)
	if err != nil {
		return
	}
	// 4.提前查询好每篇帖子的投票数
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}
	// 5.将帖子的作者以及分区信息查询出来并填充进去
	for idx, post := range posts {
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
			VoteNum:            voteData[idx],
			ApiPost:            post,
			ApiCommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}

// GetCommunityPostList 按社区获取帖子列表
func GetCommunityPostList(p *models.ParamPostList) (datas []*models.ApiPostDetail, err error) {
	// 2.按社区去redis查询id列表
	ids, err := redis.GetCommunityPostIDsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder(p) return 0 data")
		return
	}
	//根据社区id去查有哪些帖子
	posts, err := mysql.GetCommunityPostList(ids, p.CommunityID)
	datas = make([]*models.ApiPostDetail, 0)
	for _, post := range posts {
		user, err := mysql.GetUerById(post.AuthorID)
		if err != nil {
			return nil, err
		}
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			return nil, err
		}
		data := &models.ApiPostDetail{
			AuthorName:         user.Username,
			ApiPost:            post,
			ApiCommunityDetail: community,
		}
		datas = append(datas, data)
	}
	return
}

// GetPostListNew 将两个查询帖子列表的接口合二为一的函数
func GetPostListNew(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	//根据请求参数的不同实行不同的逻辑
	if p.CommunityID == 0 {
		// 查所有
		data, err = GetPostList(p)
	} else {
		// 根据社区id查询
		data, err = GetCommunityPostList(p)
	}
	if err != nil {
		zap.L().Error("GetPostListNew failed", zap.Error(err))
	}
	return
}
