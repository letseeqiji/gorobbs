package banned

import (
	"github.com/gin-gonic/gin"
	"gorobbs/model"
	"gorobbs/package/app"
	"gorobbs/package/rcode"
	"gorobbs/service/v1/user"
)

func Banned() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = rcode.SUCCESS
		ip := c.ClientIP()
		if model.CheckIpBanned(ip) {
			code = rcode.BANNED
			app.JsonErrResponse(c, code)
			c.Abort()
			return
		}

		sessions := user.GetSessions(c)
		uid := sessions.Userid
		if model.CheckUserBanned(uid) {
			code = rcode.BANNED
			app.JsonErrResponse(c, code)
			c.Abort()
			return
		}

		c.Next()
	}
}

