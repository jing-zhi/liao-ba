package logic

import (
	"liaoBa/dao/mysql"
	"liaoBa/models"
)

func GetCommunityList() ([]*models.Community, error) {
	// 查找到所有的community数据，返回
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
