package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Cross is a gin middleware function that sets the Access-Control-Allow-Origin, Access-Control-Allow-Headers, Access-Control-Allow-Methods, Access-Control-Expose-Headers, Access-Control-Max-Age, and Access-Control-Allow-Credentials headers on the response.
func Cross() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求方法
		method := c.Request.Method
		// 获取请求头中的Origin
		origin := c.Request.Header.Get("Origin")
		// 如果Origin不为空，则将*添加到Access-Control-Allow-Origin中
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			// 允许跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			// 允许跨域请求的头部
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			// 允许暴露的头部
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
			// 设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			// 允许跨域请求的凭证
			c.Header("Access-Control-Allow-Credentials", "true")
			// 设置响应的类型
			c.Set("content-type", "application/json")
		}
		// 如果请求方法为OPTIONS，则返回Options Request
		if method == "OPTIONS" {
			//c.JSON(http.StatusOK, "Options Request!")
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 调用下一个中间件
		c.Next()
	}
}
