package controller

import "github.com/gin-gonic/gin"

type BaseController struct{}

type base struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Success struct {
	base
	Data interface{} `json:"data"`
}

type RequestError struct {
	base
	Data interface{} `json:"data"`
}

type PageList struct {
	base
	Total int         `json:"total"`
	List  interface{} `json:"list"`
	Size  int         `json:"size"`
}

func (b *BaseController) SendError(ctx *gin.Context, code int8, message string, data any) {
	ctx.JSON(200, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})
}
func (b *BaseController) SendSuccess(ctx *gin.Context, code int8, message string, data any) {
	ctx.JSON(200, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})
}
func (b *BaseController) SendPage(ctx *gin.Context, code int8, message string, data any, total int64, page int64, limit int64) {
	ctx.JSON(200, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
		"total":   total,
		"page":    page,
		"limit":   limit,
	})
}
