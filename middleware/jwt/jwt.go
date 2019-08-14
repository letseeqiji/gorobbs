package jwt

import (
	"gorobbs/package/rcode"
	"gorobbs/util"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		//var data interface{}

		code = rcode.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = rcode.INVALID_PARAMS
			c.Redirect(301, "/admin/login.html")
			return
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = rcode.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = rcode.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != rcode.SUCCESS {
			c.Redirect(301, "/admin/login.html")
			return
			/*c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  rcode.GetMessage(code),
				"data": data,
			})

			c.Abort()
			return*/
		}

		c.Next()
	}
}
