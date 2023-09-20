package mysql

import (
	"reddit/models"
)

// CreatePost 创建帖子
func CreatePost(p *models.Post) (err error) {
	return db.Create(p).Error
}

// GetPostById 通过帖子id找到单个帖子内容
func GetPostById(pid int64) (post *models.ApiPost, err error) {
	post = new(models.ApiPost)
	err = db.Model(&models.Post{}).Where("post_id = ?", pid).Find(post).Error
	return
}

// GetPostList 获取帖子列表
func GetPostList(page, size int64) (posts []*models.ApiPost, err error) {
	posts = make([]*models.ApiPost, 0)
	err = db.
		Model(&models.Post{}).          //去post表中查询内容
		Offset(int((page - 1) * size)). //设置offset，跳过前几页
		Limit(int(size)).               //当前页最多显示几条
		Order("created_at DESC").       //根绝创建时间逆序排列
		Find(&posts).                   //将查询的内容放入自定义结构体Apipost中
		Error
	if err != nil {
		return nil, err
	}
	return
}

// GetPostListByIds 根据给定的id列表查询帖子数据
func GetPostListByIds(ids []string) (posts []*models.ApiPost, err error) {
	posts = make([]*models.ApiPost, 0)
	if err = db.
		Model(&models.Post{}).
		Where("post_id in ?", ids).
		Order("created_at DESC").
		Find(&posts).
		Error; err != nil {
		return nil, err
	}
	return
}

// GetCommunityPostList 获取对应社区的帖子
func GetCommunityPostList(p *models.ParamPostList) (posts []*models.ApiPost, err error) {
	posts = make([]*models.ApiPost, 0)
	err = db.
		Model(&models.Post{}).
		Where("community_id = ?", p.CommunityID).
		Order("created_at DESC").
		Find(&posts).
		Error
	if err != nil {
		return nil, err
	}
	return
}
