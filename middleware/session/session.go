package session

import (
	"github.com/gin-gonic/gin"
	"gorobbs/package/rcode"
	"gorobbs/service/v1/user"
	"net/http"
)

// 验证是否登录
func LoginCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int

		isLogin := user.IsLogin(c)

		if !isLogin {
			code = rcode.UNPASS
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  rcode.GetMessage(code),
			})
			// c.Redirect(301, "/login.html")
			c.Abort()
			return
		}

		c.Next()
	}
}
