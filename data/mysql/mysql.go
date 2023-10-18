package data

import (
	"gorm.io/gorm"
	"time"
)

func NewMysqlTable(Db *gorm.DB) {
	tables := []interface{}{
		&SysUser{},                    // 用户表
		&SysRole{},                    // 角色表
		&SysPermission{},              // 权限表
		&SysOrganization{},            // 组织机构表
		&SysUserOrganizationMapping{}, // 用户部门表
		&SysRoleResource{},            // 角色资源表
		&SysRoleLog{},                 // 角色操作日志
		&SysArticle{},                 // 文章表

	}

	err := Db.AutoMigrate(tables...)

	if err != nil {
		panic(err)
	}

	// 初始化数据
	SetDefaultValue(Db)
}

func SetDefaultValue(Db *gorm.DB) {
	admin := SysUser{
		Name:     "admin",
		Password: "E10ADC3949BA59ABBE56E057F20F883E",
		Status:   1,
		Phone:    "17608175614",
		Email:    "17608175614@163.com",
	}

	//角色表
	role := SysRole{
		Name:        "超级管理员",
		RoleKey:     "admin",
		RoleSort:    1,
		Status:      1,
		Description: "超级管理员",
	}
	//机构表
	Organization := SysOrganization{
		Name:     "xx公司",
		ParentId: 0,
		Sort:     1,
		Status:   1,
	}
	//Db.Model(&SysOrganization{}).Save(&Organization)
	//资源表
	Resource := SysRoleResource{
		ResourceName: "用户管理",
		ResourceDesc: "用户管理",
		ResourceType: 1,
		State:        1,
	}
	//Db.Model(&SysOrganization{}).Save(&Resource)
	//权限
	Permission := SysPermission{
		PermissionName: "用户管理",
		PermissionDesc: "用户管理",
	}
	//Db.Model(&SysOrganization{}).Save(&Permission)
	//部门表
	OrganizationMapping := SysUserOrganizationMapping{
		UserId:     1,
		OrganizeId: 1,
		State:      1,
		AssignTime: time.Now(),
	}
	// 存在则跳过
	Db.FirstOrCreate(&admin, admin)
	Db.FirstOrCreate(&role, role)
	Db.FirstOrCreate(&Organization, Organization)
	Db.FirstOrCreate(&Resource, Resource)
	Db.FirstOrCreate(&Permission, Permission)
	Db.FirstOrCreate(&OrganizationMapping, OrganizationMapping)
}

func Add[T SysRole | SysRoleResource | SysUser | SysPermission | SysUserOrganizationMapping | SysRoleLog](Db *gorm.DB, target T) T {
	Db.Create(&target)
	return target
}
