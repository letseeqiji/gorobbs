package v1

import (
	"gorobbs/package/app"
	"gorobbs/package/queue"
	"gorobbs/package/rcode"
	"gorobbs/package/setting"

	"github.com/gin-gonic/gin"
)

func SendRegisterMail(c *gin.Context) {
	mailTo := c.PostForm("email")
	// todo 验证邮箱合法性

	list := setting.ServerSetting.Siteurl

	err := queue.Set(list, mailTo) //发信息的操作应该丢给redis队列 然后直接返回成功给客户端
	code := rcode.SUCCESS
	if err != nil {
		code = rcode.ERROR
	}

	app.JsonOkResponse(c, code, map[string]interface{}{"mail": mailTo})
}
