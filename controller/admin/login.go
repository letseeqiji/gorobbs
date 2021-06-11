package admin

import (
	"gorobbs/model"
	"gorobbs/package/rcode"
	"gorobbs/package/session"
	"gorobbs/service/v1/user"
	"gorobbs/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminLogin(c *gin.Context) {
	sessions := user.GetSessions(c)
	c.HTML(200, "rlogin.html", gin.H{"sessions": sessions})
}

func AdminLoginCheck(c *gin.Context) {
	email := session.GetSession(c, "useremail")
	password := c.DefaultPostForm("password", "")
	code := rcode.INVALID_PARAMS

	// 2，验证邮箱和密码
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	maps["email"] = email
	user, err := model.GetUser(maps)
	if err != nil {
		code = rcode.ERROR_NOT_EXIST_USER
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": rcode.GetMessage(code),
			"data":    data,
		})
		return
	}

	// 获取加密的密码
	hashPassword := user.Password
	if !util.VerifyString(password, hashPassword) {
		code = rcode.ERROR_NOT_EXIST_USER
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": rcode.GetMessage(code),
			"data":    data,
		})
		return
	}

	/*c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  rcode.GetMessage(code),
		"data": bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(user.Password)),
	})
	return*/

	// 3，验证通过 生成token和session
	code = rcode.SUCCESS
	token, time, err := util.GenerateToken(user.Username, password)
	if err != nil {
		code = rcode.ERROR_AUTH_TOKEN
	} else {
		data["token"] = token
		data["exp_time"] = time
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": rcode.GetMessage(code),
		"data":    data,
	})
}
