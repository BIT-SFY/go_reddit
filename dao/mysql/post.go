package mysql

import "reddit/models"

// CreatePost 创建帖子
func CreatePost(p *models.Post) (err error) {
	return db.Create(p).Error
}

// GetPostById 通过帖子id找到帖子具体内容
func GetPostById(pid int64) (post *models.ApiPost, err error) {
	post = new(models.ApiPost)
	err = db.Model(&models.Post{}).Where("post_id = ?", pid).Find(post).Error
	return
}

// GetPostList 获取帖子列表
func GetPostList(page, size int64) (posts []*models.ApiPost, err error) {
	posts = make([]*models.ApiPost, 0)
	if err = db.Model(&models.Post{}).Offset(int((page - 1) * size)).Limit(int(size)).Find(&posts).Error; err != nil {
		return nil, err
	}
	return
}
