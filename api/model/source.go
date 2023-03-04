package model

type Source struct {
	Base
	Author    string `json:"author" gorm:"column:AUTHOR"`
	Key       string `json:"key" gorm:"column:KEY"`
	Name      string `json:"name" gorm:"column:NAME"`
	Url       string `json:"url" gorm:"column:URL"`
	IsChecked bool   `json:"is_checked"`
}
