package data

import (
	"gorm.io/gorm"
	"time"
)

// SysRole 角色
type SysRole struct {
	gorm.Model
	Name        string `gorm:"type:varchar(20);not null; comment: '角色名称'"`
	RoleKey     string `gorm:"type:varchar(20);not null; comment: '角色权限字符串'"`
	RoleSort    int    `gorm:"type:int;not null; comment: '显示顺序'"`
	Status      int    `gorm:"type:tinyint; default:1; comment: '0:禁用 1:启用'"`
	Description string `gorm:"type:varchar(255); comment: '描述'"`
}

// SysPermission 权限表
type SysPermission struct {
	gorm.Model
	PermissionName string `gorm:"type:varchar(20);not null;comment: '权限名称'"`
	PermissionDesc string `gorm:"type:varchar(20);not null;comment: '权限字符串'"`
}

// SysOrganization 组织机构表
type SysOrganization struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null; comment: '组织名称'"`
	ParentId uint   `gorm:"type:int; comment: '父级ID'"`
	Sort     int    `gorm:"type:int;not null; comment: '排序'"`
	Status   int    `gorm:"type:tinyint; default:1; comment: '0:禁用 1:启用'"`
}

// SysUserOrganizationMapping 用户部门表
type SysUserOrganizationMapping struct {
	gorm.Model
	UserId     uint      `gorm:"type:int;not null;references: sys_user(id);comment: '用户ID'"`
	OrganizeId uint      `gorm:"type:int;not null;references: sys_organization(id); comment: '组织ID'"`
	State      int       `gorm:"type:tinyint; default:1; comment: '0:禁用 1:启用'"`
	AssignTime time.Time `gorm:"type:datetime; comment: '分配时间'"`
}

// SysRoleResource  角色资源表
type SysRoleResource struct {
	gorm.Model
	ResourceName string `gorm:"type:varchar(20);not null;comment: '资源名称'"`
	ResourceDesc string `gorm:"type:varchar(20);not null;comment: '资源描述'"`
	ResourceType int    `gorm:"type:tinyint; not null; comment: '资源类型 1:菜单 2:按钮'"`
	State        int    `gorm:"type:tinyint; default:1; comment: '0:禁用 1:启用'"`
}

// SysRoleLog  角色操作日志
type SysRoleLog struct {
	gorm.Model
	OperateTime     time.Time `gorm:"type:datetime; comment: '操作时间'"`
	OperateUser     string    `gorm:"type:varchar(20); references:sys_user(id); comment: '操作人'"`
	OperateContent  string    `gorm:"type:varchar(20); comment: '操作内容'"`
	OperateType     int       `gorm:"type:tinyint; comment: '操作类型 1:新增 2:修改 3:删除'"`
	OperateResource string    `gorm:"type:varchar(20);references:sys_role_resource(id); comment: '操作资源'"`
	OperateResult   int       `gorm:"type:tinyint; comment: '操作结果 1:成功 2:失败'"`
	OperateIp       string    `gorm:"type:varchar(20); comment: '操作IP'"`
	Detail          string    `gorm:"type:text; comment: '详细信息'"`
}

func (SysRole) TableName() string {
	return "sys_role"
}

func (SysPermission) TableName() string {
	return "sys_permission"
}

func (SysOrganization) TableName() string {
	return "sys_organization"
}

func (SysUserOrganizationMapping) TableName() string {
	return "sys_user_organization_mapping"
}

func (SysRoleResource) TableName() string {
	return "sys_role_resource"
}

func (SysRoleLog) TableName() string {
	return "sys_role_log"
}
