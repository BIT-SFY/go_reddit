package mysql

import "reddit/models"

// CreatePost 创建帖子
func CreatePost(p *models.Post) (err error) {
	return db.Create(p).Error
}

// GetPostById 通过帖子id找到帖子具体内容
func GetPostById(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	err = db.Where("post_id = ?", pid).Take(post).Error
	return
}
