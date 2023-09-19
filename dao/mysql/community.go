package mysql

import (
	"reddit/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// GetCommunityList 获取社区列表
func GetCommunityList() (data []*models.ApiCommunityDetail, err error) {
	if err = db.Model(&models.Community{}).Find(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}

// GetCommunityDetailByID 通过社区ID获取社区信息
func GetCommunityDetailByID(id int64) (data *models.ApiCommunityDetail, err error) {
	data = new(models.ApiCommunityDetail)
	if err = db.Model(&models.Community{}).Where("community_id = ?", id).Take(data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.L().Warn("communityID is not correct")
			err = ErrorInvalidID
		}
	}
	return
}
