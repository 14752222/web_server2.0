package middleware

import (
	"github.com/gin-gonic/gin"
	"web_server_2.0/bootstrap"
	"web_server_2.0/utils"
)

func AuthMiddleware(app *bootstrap.Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取请求头中的token
		token := c.Request.Header.Get("authorization")
		if token == "" {
			c.JSON(401, gin.H{
				"code": 401,
				"msg":  "请求未携带token，无权限访问",
			})
			c.Abort()
		}
		token = token[7:]

		secret := app.Env.SecretKey
		//解析token
		_, err := utils.IsAuthorized(token, secret)
		if err != nil {
			c.JSON(401, gin.H{
				"code": 401,
				"msg":  "无效的token",
			})
			c.Abort()
		}

		userID, err := utils.ExtractIDFromToken(token, secret)

		if err != nil {
			c.JSON(401, gin.H{
				"code": 401,
				"msg":  "无效的token",
			})
			c.Abort()
		}
		c.Set("x-user-id", userID)
		c.Next()
	}
}

/**
	拦截重复请求ip中间件
**/

func RepeatRequestMiddleware(app *bootstrap.Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		//	当前ip是否重复请求
		ip := c.ClientIP()
		requestIp := utils.RequestInfo{
			Ip:    ip,
			Api:   c.Request.URL.Path,
			Redis: app.Redis,
			Ctx:   c.Request.Context(),
		}
		if requestIp.IsRepeat() {
			c.JSON(401, gin.H{
				"code": 401,
				"msg":  "请求过于频繁，请稍后再试",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
