package service

import (
	"tracker-manager-go/conf"
	"tracker-manager-go/model"
)

type SourceService struct {
}

func (SourceService) List() []*model.Source {
	var sources []*model.Source
	conf.DB.Model(&model.Source{}).Order("create_time desc").Find(&sources)
	return sources
}

func (SourceService) UpdateById(source model.Source) bool {
	if e := conf.DB.Model(&model.Source{}).Where("uuid = ?", source.UUID).Updates(source); e != nil {
		return false
	}
	return true
}
