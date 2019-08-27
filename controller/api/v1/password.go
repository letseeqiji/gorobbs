package v1

import (
	"gorobbs/model"
	"gorobbs/package/app"
	"gorobbs/package/rcode"
	"gorobbs/package/setting"
	email_service "gorobbs/service/v1/email"
	"gorobbs/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SendResetPasswordEmail(c *gin.Context) {
	mailTo := c.PostForm("email")
	// todo 验证邮箱合法性
	port := strconv.Itoa(setting.ServerSetting.HttpPort)
	host := setting.ServerSetting.Sitepre + "://" + setting.ServerSetting.Siteurl
	if setting.ServerSetting.RunMode == "debug" {
		host = host + ":" + port
	}

	err := email_service.SendResetPasswordMail(host, mailTo) //发信息的操作应该丢给redis队列 然后直接返回成功给客户端
	code := rcode.SUCCESS
	if err != nil {
		code = rcode.ERROR
	}

	app.JsonOkResponse(c, code, map[string]interface{}{"mail": mailTo})
}

// 用户重设密码
func UserResetPassword(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	code := rcode.SUCCESS

	password, _ = util.BcryptString(password)

	var wmap = make(map[string]interface{})
	wmap["email"] = email
	err := model.UpdateUser(wmap, map[string]interface{}{"password": password})
	if err != nil {
		code = rcode.ERROR_SQL_UPDATE_FAIL
		app.JsonErrResponse(c, code)
		return
	}

	app.JsonOkResponse(c, code, nil)
}