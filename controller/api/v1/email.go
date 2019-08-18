package v1

import (
	"gorobbs/package/rcode"
	email_service "gorobbs/service/v1/email"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendRegisterMail(c *gin.Context) {
	mailTo := c.PostForm("email")
	// todo 验证邮箱合法性
	host := c.Request.Host

	err := email_service.SendRegisterMail2(host, mailTo) //发信息的操作应该丢给redis队列 然后直接返回成功给客户端
	code := rcode.SUCCESS
	if err != nil {
		code = rcode.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"message":  err.Error(),
			"data": map[string]interface{}{"mail":mailTo},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"message":  rcode.GetMessage(code),
		"data": map[string]interface{}{"mail":mailTo},
	})
}
