package data

import (
	"gorm.io/gorm"
)

type SysArticle struct {
	gorm.Model
	Uuid     string `gorm:"type:varchar(100);not null;uniqueIndex"`
	Title    string `gorm:"type:varchar(100);not null;"`
	Password string `gorm:"type:varchar(100);not null;"`
	Account  string `gorm:"type:varchar(100);not null;"`
	//IsCheck  int    `gorm:"type:tinyint default:0;"`
	IsCheck int `gorm:"type:tinyint; default:0; comment: '0:禁用 1:启用'"`
}

func (SysArticle) TableName() string {
	return "sys_articles"
}
