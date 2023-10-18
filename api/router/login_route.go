package router

import (
	"github.com/gin-gonic/gin"
	"web_server_2.0/api/controller"
	"web_server_2.0/api/repository"
	"web_server_2.0/bootstrap"
)

func NewLoginRoute(app *bootstrap.Application, router *gin.RouterGroup) {
	control := controller.LoginController{
		Env:        app.Env,
		Db:         app.Mysql,
		Repository: &repository.LoginRepository{},
		Redis:      app.Redis,
	}
	router.POST("/login", control.Login)
	router.POST("/register", control.Register)
	router.POST("/sendEmail", control.SendEmail)
	router.POST("/auth_code_login", control.AuthCodeLogin)
}
