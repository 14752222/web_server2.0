package repository

import (
	"fmt"
	"gorm.io/gorm"
	data "web_server_2.0/data/mysql"
)

type LoginRepository struct{}

// IsUserExist
// 用户是否存在
func (lr *LoginRepository) IsUserExist(db *gorm.DB, email string) (bool, data.SysUser) {
	var user data.SysUser
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return false, user
	}
	return true, user
}

func (lr *LoginRepository) CreateUser(db *gorm.DB, email string, password string) (bool, data.SysUser) {
	var user data.SysUser
	user.Email = email
	user.Password = password
	err := db.Create(&user).Error
	if err != nil {
		fmt.Println(`创建用户失败 err`, err)
		return false, user
	}
	return true, user
}

// AddUserPermission 添加用户权限
func (lr *LoginRepository) AddUserPermission(db *gorm.DB, userId int, permissionId int) (bool, data.SysUser) {
	var user data.SysUser
	err := db.Model(&user).Where("id = ?", userId).Update("permission_id", permissionId).Error
	if err != nil {
		return false, user
	}
	return true, user
}

// GetUserPermission 获取用户权限
func (lr *LoginRepository) GetUserPermission(db *gorm.DB, userId int) (bool, data.SysUser) {
	var user data.SysUser
	err := db.Model(&user).Where("id = ?", userId).Find(&user).Error
	if err != nil {
		return false, user
	}
	return true, user
}

// UpdateUserPassword 修改用户密码
func (lr *LoginRepository) UpdateUserPassword(db *gorm.DB, userId int, password string) (bool, data.SysUser) {
	var user data.SysUser
	err := db.Model(&user).Where("id = ?", userId).Update("password", password).Error
	if err != nil {
		return false, user
	}
	return true, user
}
