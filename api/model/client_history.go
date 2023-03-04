package model

type ClientHistory struct {
	Base
	ClientId       string `json:"client_id" gorm:"column:CLIENT_ID"`
	SyncSourceId   string `json:"sync_source_id" gorm:"column:SYNC_SOURCE_ID"`
	TrackerContent string `json:"tracker_content" gorm:"column:TRACKER_CONTENT"`
}
