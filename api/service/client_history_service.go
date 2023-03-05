package service

import "tracker-manager-go/model"

type ClientHistoryService struct {
}

type ClientHistoryVO struct {
	model.ClientHistory
	ClientName string `json:"client_name" gorm:"column:client_name"`
}

func (ClientHistoryService) List() []*ClientHistoryVO {
	var vos []*ClientHistoryVO
	db.Model(&model.ClientHistory{}).Select(
		"*, (select name from client c where c.uuid = client_id) as client_name").Find(&vos)
	return vos
}

func (s ClientHistoryService) DeleteByIds(ids []string) error {
	return db.Where("uuid in ?", ids).Delete(&model.ClientHistory{}).Error
}
