package model

type SyncLog struct {
	Base
	SourceId   string `json:"source_id"`
	SourceName string `json:"source_name"`
	SuccessNum uint   `json:"success_num"`
	ErrorNum   uint   `json:"error_num"`
}
