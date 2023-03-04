package service

import (
	"gorm.io/gorm"
	"sync"
	"tracker-manager-go/model"
	"tracker-manager-go/util"
)

type SyncService struct {
}

type SyncSourceDTO struct {
	Source model.Source
	Value  string
}

var wg sync.WaitGroup

func getTracker(sources []*model.Source) []*SyncSourceDTO {
	dto := make([]*SyncSourceDTO, 0)
	for _, source := range sources {
		wg.Add(1)
		go func(s *model.Source) {
			defer wg.Done()
			body, _ := util.Get(s.Url)
			dto = append(dto, &SyncSourceDTO{
				Source: *s,
				Value:  body,
			})
		}(source)
	}
	wg.Wait()
	return dto
}

func (SyncService) Sync() error {
	var sources []*model.Source
	db.Model(&model.Source{}).Find(&sources)
	dtos := getTracker(sources)
	l := len(dtos)
	syncSources := make([]*model.SyncSource, 0)
	ids := ""
	names := ""
	for index, dto := range dtos {
		syncSources = append(syncSources, &model.SyncSource{
			SourceId: dto.Source.UUID,
			Content:  dto.Value,
		})
		ids += dto.Source.UUID
		names += dto.Source.Name
		if index < l-1 {
			ids += ","
			names += ","
		}
	}
	err := db.Transaction(func(tx *gorm.DB) error {
		// 删除原数据
		tx.Model(&model.SyncSource{}).Where("1 = 1").Delete(&model.SyncSource{})
		// 写入新数据
		if e := tx.Create(&syncSources).Error; e != nil {
			return e
		}
		// 记录日志
		syncLog := model.SyncLog{
			SourceId:   ids,
			SourceName: names,
			SuccessNum: uint(l),
			ErrorNum:   uint(len(sources) - l),
		}
		if e := tx.Model(&model.SyncLog{}).Create(&syncLog).Error; e != nil {
			return e
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (SyncService) Clear() error {
	return db.Model(&model.SyncSource{}).Where("1 = 1").Delete(&model.SyncSource{}).Error
}
