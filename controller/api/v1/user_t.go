package v1

import (
	"gorobbs/package/session"

	"github.com/gin-gonic/gin"
)

func TestSetSesssion(c *gin.Context) {
	session.SetSession(c, "username", "zhanglei")
	c.JSON(200, gin.H{"msg": "ok"})
}

func TestGetSesssion(c *gin.Context) {
	u := session.GetSession(c, "username")
	if len(u) > 0 {
		c.JSON(200, gin.H{"msg": "ok 不为空", "user": u, "ulen": len(u)})
	} else {
		c.JSON(200, gin.H{"msg": "ok", "user": u, "ulen": len(u)})
	}

}

func TestDelSesssion(c *gin.Context) {
	session.DeleteSession(c, "username")

	c.JSON(200, gin.H{"msg": "ok"})
}

func TestEncyPasscheck(c *gin.Context) {

}
