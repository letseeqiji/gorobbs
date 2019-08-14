package xss

import (
	"github.com/gin-gonic/gin"
	package_redis "gorobbs/package/gredis"
)

func XSS() gin.HandlerFunc {
	return func(c *gin.Context) {
		xssToken := c.DefaultPostForm("xss_token", "")
		if len(xssToken)==0 {
			c.JSON(200, gin.H{
				"code": 401,
				"message":  "请提交xsstoken",
			})
			c.Abort()
			return
		}
		_, err := package_redis.Get(xssToken)
		if err == nil {
			c.JSON(200, gin.H{
				"code": 403,
				"message":  "已经提交过了，不要重复提交",
			})
			c.Abort()
			return
		}
		package_redis.Set(xssToken, xssToken, 100)

		c.Next()
	}
}
