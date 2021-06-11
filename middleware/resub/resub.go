package resub

import (
	"github.com/gin-gonic/gin"
	package_redis "gorobbs/package/gredis"
)

func RESUB() gin.HandlerFunc {
	return func(c *gin.Context) {
		uniqueToken := c.DefaultPostForm("unique_token", "")
		if len(uniqueToken)==0 {
			c.JSON(200, gin.H{
				"code": 401,
				"message":  "非法请求",
			})
			c.Abort()
			return
		}
		_, err := package_redis.Get(uniqueToken)
		if err == nil {
			c.JSON(200, gin.H{
				"code": 403,
				"message":  "已经提交过了，不要重复提交",
			})
			c.Abort()
			return
		}
		package_redis.Set(uniqueToken, uniqueToken, 100)

		c.Next()
	}
}
