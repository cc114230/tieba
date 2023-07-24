package logic

import (
	"tieba/dao/mysql"
	"tieba/models"
)

func GetCommunityList() ([]*models.Community, error) {

	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
