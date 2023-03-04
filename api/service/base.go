package service

import (
	"tracker-manager-go/conf"
)

var db = conf.DB

type Base[T any] struct {
}

func (Base[T]) GetById(id string) *T {
	var t T
	db.Model(&t).Where("uuid = ?", id).First(&t)
	return &t
}

func (Base[T]) UpdateById(id string, params map[string]any) error {
	var model T
	return db.Model(&model).Where("uuid = ?", id).Updates(params).Error
}

func (Base[T]) List() []*T {
	var model T
	var ts []*T
	db.Model(&model).Order("create_time desc").Find(&ts)
	return ts
}

func (Base[T]) Save(t T) error {
	return db.Create(t).Error
}

func (Base[T]) DeleteById(id string) error {
	var model T
	return db.Model(&model).Where("uuid = ?", id).Delete(&model).Error
}
