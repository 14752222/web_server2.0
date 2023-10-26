package controller

import "github.com/gin-gonic/gin"

type BaseController struct{}

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
