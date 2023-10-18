package middleware

import "github.com/gin-gonic/gin"

// 查询用户权限是否存在

func QueryUserPermission() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取用户权限
		//获取c.Set("x-user-id")的值
		//userId, ok := c.Get("x-user-id")
		//if !ok {
		//	c.JSON(401, gin.H{
		//		"code": 401,
		//		"msg":  "无效的token",
		//	})
		//	c.Abort()
		//}
		// 查询用户权限是否存在

		// 存在则继续
		// 不存在则返回错误
	}
}
