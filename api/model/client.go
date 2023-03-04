package model

type Client struct {
	Base
	Name     string `json:"name" gorm:"column:NAME"`
	Url      string `json:"url" gorm:"column:URL"`
	Username string `json:"username" gorm:"column:USERNAME"`
	Password string `json:"password" gorm:"column:PASSWORD"`
}
