package model

type SyncSource struct {
	Base
	SourceId string `json:"source_id"`
	Content  string `json:"content"`
}
