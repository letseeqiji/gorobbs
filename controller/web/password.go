package web

import (
	email_package "gorobbs/package/email"
	"gorobbs/package/setting"
	"gorobbs/service/v1/user"
	"gorobbs/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ForgetPassword(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)
	// 网站描述
	webname := setting.ServerSetting.Sitename
	description := setting.ServerSetting.Sitebrief
	forumname := "忘记密码"

	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"forgetpass.html",
		// Pass the data that the page uses
		gin.H{
			"title":       "用户重置密码",
			"islogin":     islogin,
			"sessions":    sessions,
			"forums":      forums,
			"webname":     webname,
			"description": description,
			"forumname":   forumname,
		},
	)
}

func SendEmailOk(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)
	// 网站描述
	webname := setting.ServerSetting.Sitename
	description := setting.ServerSetting.Sitebrief
	forumname := "发送成功"

	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"send_reset_email_ok.html",
		// Pass the data that the page uses
		gin.H{
			"title":       "用户重置密码",
			"islogin":     islogin,
			"sessions":    sessions,
			"forums":      forums,
			"webname":     webname,
			"description": description,
			"forumname":   forumname,
		},
	)
}

func ResetForgetPassword(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)
	// 网站描述
	webname := setting.ServerSetting.Sitename
	description := setting.ServerSetting.Sitebrief
	forumname := "重设密码"

	email := c.DefaultQuery("email", "")
	time := c.DefaultQuery("time", "")
	sign := c.DefaultQuery("sign", "")

	// 链接地址错误
	if sign != util.EncodeMD5(email+time) {
		c.HTML(
			http.StatusOK,
			"reset_email_err.html",
			gin.H{
				"title":       "重设密码",
				"islogin":     islogin,
				"sessions":    sessions,
				"forums":      forums,
				"webname":     webname,
				"description": description,
				"forumname":   forumname,
			})
		return
	}

	// quredis中获取邮箱，如果能获取到 超时验证 需要冲洗注册提交
	if _, err := email_package.Get(email); err != nil {
		c.HTML(
			http.StatusOK,
			"reset_email_pass_time.html",
			gin.H{
				"title":       "重设密码",
				"islogin":     islogin,
				"sessions":    sessions,
				"forums":      forums,
				"webname":     webname,
				"description": description,
				"forumname":   forumname,
			})
		return
	}

	// 更新用户邮箱验证状态为以已经验证并删除redis中该条数据
	email_package.Delete(email)

	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"reset_password.html",
		// Pass the data that the page uses
		gin.H{
			"title":       "重设密码",
			"islogin":     islogin,
			"sessions":    sessions,
			"forums":      forums,
			"webname":     webname,
			"description": description,
			"forumname":   forumname,
			"email":       email,
		},
	)
}