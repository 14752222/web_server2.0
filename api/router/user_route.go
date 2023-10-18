package router

import (
	"github.com/gin-gonic/gin"
	"web_server_2.0/api/controller"
	"web_server_2.0/bootstrap"
)

func NewUserRoute(app *bootstrap.Application, router *gin.RouterGroup) {
	control := controller.UserController{Env: app.Env, Db: app.Mysql, Redis: app.Redis}
	router.POST("/getUser", control.GetUserInfo)
	router.POST("/getUserInfo", control.GetUserInfo)
	router.POST("/getPermission ", control.GetUserPermission)
	router.POST("/get_user_info", control.GetUserInfo)

}
