package model

type SyncSource struct {
	Base
	SourceId string `json:"source_id" gorm:"column:SOURCE_ID"`
	Content  string `json:"content" gorm:"column:CONTENT"`
}
