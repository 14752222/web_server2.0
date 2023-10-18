package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"web_server_2.0/config"
)

type UserController struct {
	Env   *config.Env
	Db    *gorm.DB
	Redis *redis.Client
}

// GetUserInfo 获取用户信息
// @Summary 获取用户信息
// @Tags 用户
// @Accept  json
// @Produce  json
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /api/v1/login/get_user_info [post]
func (uc *UserController) GetUserInfo(ctx *gin.Context) {
	//获取token
	user, err := ctx.Get("x-user-id")
	if !err {
		ctx.JSON(200, gin.H{
			"code": 400,
			"msg":  "获取用户信息失败",
			"data": nil,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "获取用户信息成功",
		"data": user,
	})
}

// GetUserPermission 获取用户权限
// @Summary 获取用户权限
// @Tags 用户
// @Accept  json
// @Produce  json
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /api/v1/login/get_UserInfo [post]
func (uc *UserController) GetUserPermission(ctx *gin.Context) {
	//获取token
	user, err := ctx.Get("x-user-id")
	if !err {
		ctx.JSON(200, gin.H{
			"code": 400,
			"msg":  "获取用户信息失败",
			"data": nil,
		})
		return
	}
	fmt.Println(user)
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "获取用户权限成功",
		"data": user,
	})

	//获 user 权限列表

}
