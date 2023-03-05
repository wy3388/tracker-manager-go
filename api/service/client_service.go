package service

import (
	"errors"
	"strings"
	"tracker-manager-go/client"
	"tracker-manager-go/model"
)

type ClientService struct {
	Base[model.Client]
}

func (cs ClientService) Sync(id string) error {
	c := cs.GetById(id)
	if c.Url != "" && c.Username != "" && c.Password != "" {
		var syncSource []*model.SyncSource
		if err := db.Model(&model.SyncSource{}).Where("source_id in (?)",
			db.Model(&model.Source{}).Select("uuid").Where("is_checked = ?",
				1)).Find(&syncSource).Error; err != nil {
			return err
		}
		if len(syncSource) == 0 {
			return errors.New("请先同步Tracker")
		}
		text := ""
		ids := ""
		for _, source := range syncSource {
			content := source.Content
			if strings.Contains(content, "\n\n") {
				content = strings.ReplaceAll(content, "\n\n", "\n")
			}
			ids += source.UUID
			ids += ","
			text += content
			text += "\n"
		}
		if c.Name == "qBittorrent" {
			err := client.Sync(&client.QBittorrent{}, *c, text)
			if err != nil {
				return err
			}
			// 写入日志
			db.Model(&model.ClientHistory{}).Create(&model.ClientHistory{
				ClientId:       id,
				SyncSourceId:   ids[:len(ids)-1],
				TrackerContent: text,
			})
			return nil
		}
		return errors.New("客户端暂未支持")
	}
	return errors.New("记录不存在")
}
