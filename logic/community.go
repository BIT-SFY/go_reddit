package logic

import (
	"reddit/dao/mysql"
	"reddit/models"
)

// GetCommunityList 查询所有社区
func GetCommunityList() ([]*models.ApiCommunityDetail, error) {
	//查数据库 查找到所有的community并返回
	return mysql.GetCommunityList()
}

// GetCommunityListDetail 根据ID查询社区详情
func GetCommunityListDetail(id int64) (*models.ApiCommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
