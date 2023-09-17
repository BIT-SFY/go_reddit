package mysql

import (
	"reddit/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func GetCommunityList() (data []*models.Community, err error) {
	if err = db.Select("community_id,community_name,introduction").Where("").Find(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}

// GetCommunityDetailByID 通过社区ID获取社区信息
func GetCommunityDetailByID(id int64) (data *models.Community, err error) {
	data = new(models.Community)
	if err = db.Where("community_id = ?", id).Take(data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.L().Warn("communityID is not correct")
			err = ErrorInvalidID
		}
	}
	return
}
