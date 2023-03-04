package model

type SyncLog struct {
	Base
	SourceId   string `json:"source_id" gorm:"column:SOURCE_ID"`
	SourceName string `json:"source_name" gorm:"column:SOURCE_NAME"`
	SuccessNum uint   `json:"success_num" gorm:"column:success_num"`
	ErrorNum   uint   `json:"error_num" gorm:"column:error_num"`
}
