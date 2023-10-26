package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"time"
	"web_server_2.0/api/repository"
	repository2 "web_server_2.0/api/repository"
	"web_server_2.0/client"
	"web_server_2.0/config"
	"web_server_2.0/utils"
)

type ArticleController struct {
	Result     *BaseController
	Env        *config.Env
	Db         *gorm.DB
	Repository *repository.ArticleRepository
	Redis      *redis.Client
	Minio      *client.FileManager
}

// AddArticle 添加文章
// @Summary 添加文章
// @Description 添加文章
// @Tags 文章
// @Accept  json
// @Produce  json
// @Param article body repository2.Article true "文章"
// @Success 200 {object} string "success"
// @Router /user/article/add [post]
// @Security ApiKeyAuth
// @Failure 400 {object} string "error"
func (c *ArticleController) AddArticle(ctx *gin.Context) {
	// 获取参数
	var article repository2.Article
	err := ctx.ShouldBind(&article)

	if err != nil {
		c.Result.SendError(ctx, 0, "参数错误", nil)
		return
	}

	// 保存到数据库
	uuid := article.AddArticle(c.Db)
	htmloptions := &utils.HtmlOptions{
		Title:    article.Title,
		Body:     article.Content,
		FileName: fmt.Sprintf(`tmp/%s.html`, uuid),
	}
	ok := htmloptions.CreateHtml()

	if !ok {
		c.Result.SendError(ctx, -1, "系统错误", nil)
		return
	}
	fileInfo, err := c.Minio.UploadFile(&client.FileDesc{
		Bucket:      "articlehtml",
		FileName:    fmt.Sprintf(`%s.html`, uuid),
		FilePath:    htmloptions.FileName,
		ContentType: "application/html",
	})
	if err != nil {
		c.Result.SendError(ctx, -1, "系统错误", nil)
		return
	}

	url := "/html/article/" + uuid
	qr := utils.CreateQrImage(url)

	data := map[string]interface{}{
		"url":      url,
		"qr":       qr,
		"location": fileInfo.Location,
	}

	// 返回结果
	c.Result.SendSuccess(ctx, 1, "成功", data)

	htmloptions.DeleteFile()

	//第七天执行删除任务 c.Minio.DeleteFile(fileInfo.Bucket, fileInfo.Key)
	tanget := time.Now().Add(time.Hour * 24 * 7)
	//tanget := time.Now().Add(time.Second * 10)
	time.AfterFunc(tanget.Sub(
		time.Now()),
		func() {
			article.DeleteArticle(c.Db, uuid)
			err := c.Minio.RemoveFile(fileInfo.Bucket, fileInfo.Key)
			if err != nil {
				fmt.Println(`删除文件失败`)
				return
			}
		})
	//cronId, err := utils.AddCronJob(fmt.Sprintf(`0 0 0 %v * *`, tanget), func() {
	//	err := c.Minio.RemoveFile(fileInfo.Bucket, fileInfo.Key)
	//	if err != nil {
	//		fmt.Println(`删除文件失败`)
	//		return
	//	}
	//})
	//if err != nil {
	//	fmt.Println(`添加定时任务失败`, err)
	//	return
	//}
	//fmt.Println(`添加定时任务成功`, cronId)
	//c.Redis.Set(ctx, fmt.Sprintf(`cron:%s`, cronId), cronId, tanget.Sub(time.Now()))
}

// GetArticleList 获取文章列表
// @Summary 获取文章列表
// @Description 获取文章列表
// @Tags 文章
// @Accept  json
// @Produce  json
// @Param page query int true "页码"

func (c *ArticleController) GetArticle(ctx *gin.Context) {
	// 读取/article/:id
	articleId := ctx.Param("id")
	article := c.Repository.GetArticle(c.Db, articleId)
	if article == nil {
		ctx.HTML(200, "404.html", gin.H{})
		return
	}
	//是否需要密码
	if article.IsCheck == 1 {
		//重定向
		ctx.Redirect(302, fmt.Sprintln("/article/"+articleId+"/password"))
		return
	}
	fmt.Println(`title`, article.Title, "content", article.Content)
	ctx.HTML(200, "article/index.html", gin.H{
		"title":   article.Title,
		"content": article.Content,
	})
}

func (c *ArticleController) GetArticleList(ctx *gin.Context) {
	// 读取/article/:id
	//articleList := c.Repository.GetArticleList(c.Db)
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": []string{"1", "2", "3"},
	})

}
