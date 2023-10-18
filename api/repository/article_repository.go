package repository

import (
	"gorm.io/gorm"
	"strings"
	data "web_server_2.0/data/mysql"
	"web_server_2.0/utils"
)

type ArticleRepository struct{}

type Article struct {
	Title    string `json:"title" form:"title" binding:"required"`
	Content  string `json:"content" form:"content" binding:"required"`
	Account  string `json:"account" form:"account"`
	Password string `json:"password" form:"password"`
	IsCheck  int    `json:"isCheck"`
	uuid     string
}

func (a *Article) AddArticle(db *gorm.DB) string {
	Uuid := utils.CreateUUID()
	Uuid = strings.ReplaceAll(Uuid, "-", "")
	a.uuid = Uuid

	article := data.SysArticle{
		Uuid:     Uuid,
		Title:    a.Title,
		Account:  a.Account,
		Password: a.Password,
	}
	db.Create(&article)

	return Uuid
}

func (a *Article) DeleteArticle(db *gorm.DB, uuid string) {
	if uuid == "" {
		uuid = a.uuid
	}

	//查询是否
	article := data.SysArticle{
		Uuid: uuid,
	}
	db.First(&article, "uuid = ?", uuid)
	db.Delete(&article)
}

func (a *ArticleRepository) GetArticle(db *gorm.DB, id string) *Article {
	//查询是否
	article := data.SysArticle{
		Uuid: id,
	}
	db.First(&article, "uuid = ?", id)

	return &Article{
		Title:    article.Title,
		IsCheck:  article.IsCheck,
		Account:  article.Account,
		Password: article.Password,
	}
}
