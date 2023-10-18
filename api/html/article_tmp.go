package html

import (
	"github.com/gin-gonic/gin"
	"web_server_2.0/api/controller"
	"web_server_2.0/bootstrap"
)

func NewArticleTmp(app *bootstrap.Application, router *gin.RouterGroup) {
	control := controller.ArticleController{Env: app.Env, Db: app.Mysql, Redis: app.Redis, Minio: app.Minio}
	// 文章模板
	router.GET("/article/:id", control.GetArticle)

}
