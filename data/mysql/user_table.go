package data

import "gorm.io/gorm"

type SysUser struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null; comment: '姓名'"`
	WxOpenid string `gorm:"type:varchar(50);not null; comment: '微信openid'"`
	Password string `gorm:"type:varchar(32);not null; comment: '密码'"`
	Email    string `gorm:"type:varchar(50);not null; unique comment: '邮箱'"`
	Phone    string `gorm:"type:varchar(20);not null; unique comment: '手机号'"`
	Gender   string `gorm:"type:tinyint; default:0 ; comment: '0:未知 1:男 2:女'"`
	Avatar   string `gorm:"type:varchar(255); comment: '头像'"`
	Status   int    `gorm:"type:tinyint; default:1; comment: '0:禁用 1:启用'"`
}

func (SysUser) TableName() string {
	return "sys_user"
}
