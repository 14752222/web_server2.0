package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"web_server_2.0/config"
)

type UserController struct {
	Result *BaseController
	Env    *config.Env
	Db     *gorm.DB
	Redis  *redis.Client
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
		uc.Result.SendError(ctx, -1, "获取用户信息失败", nil)
		return
	}

	uc.Result.SendSuccess(ctx, 200, "获取用户信息成功", user)
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
		//ctx.JSON(200, gin.H{
		//	"code": 400,
		//	"msg":  "获取用户信息失败",
		//	"data": nil,
		//})
		uc.Result.SendError(ctx, -1, "获取用户信息失败", nil)
		return
	}
	//fmt.Println(user)
	//ctx.JSON(200, gin.H{
	//	"code": 200,
	//	"msg":  "获取用户权限成功",
	//	"data": user,
	//})
	uc.Result.SendSuccess(ctx, 200, "获取用户权限成功", user)
}
