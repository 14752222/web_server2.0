package router

import (
	"github.com/gin-gonic/gin"
	"web_server_2.0/api/controller"
	"web_server_2.0/api/repository"
	"web_server_2.0/bootstrap"
)

func NewUserArticle(app *bootstrap.Application, router *gin.RouterGroup) {
	control := controller.ArticleController{Env: app.Env, Db: app.Mysql, Redis: app.Redis, Repository: &repository.ArticleRepository{}, Minio: app.Minio}
	router.POST("/user/article/add", control.AddArticle)
	router.GET("/user/article/list", control.GetArticleList)
}
