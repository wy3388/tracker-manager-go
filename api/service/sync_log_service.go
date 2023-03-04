package service

import "tracker-manager-go/model"

type SyncLogService struct {
	Base[model.SyncLog]
}

func (SyncLogService) DeleteByIds(ids []string) error {
	return db.Where("uuid in ?", ids).Delete(&model.SyncLog{}).Error
}
