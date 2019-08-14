package online

import (
	"github.com/gin-gonic/gin"
	package_online "gorobbs/package/online"
)

func OnLine() gin.HandlerFunc {
	return func(c *gin.Context) {
		uip := c.ClientIP()
		// 设置
		package_online.Set("online:"+uip, 1, 15*60)

		c.Next()
	}
}