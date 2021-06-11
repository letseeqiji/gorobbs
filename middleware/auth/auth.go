package auth

import (
	"github.com/gin-gonic/gin"
	"gorobbs/package/session"
)

func AUTH() gin.HandlerFunc {
	return func(c *gin.Context) {
		if session.GetSession(c, "isadmin") != "1" {
			c.Redirect(301, "/index.html")
			return
		}
		c.Next()
	}
}
