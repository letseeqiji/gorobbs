package web

import (
	email_package "gorobbs/package/email"
	"gorobbs/package/session"
	"gorobbs/package/setting"
	"gorobbs/service/v1/user"
	"gorobbs/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*var (
	forums []model.Forum
)*/

/*func init() {
	forums = layout_service.GetForumList()
}*/

func Register(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)
	// 网站描述
	webname := setting.ServerSetting.Sitename
	description := setting.ServerSetting.Sitebrief
	forumname := "注册"

	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"register.html",
		// Pass the data that the page uses
		gin.H{
			"title":       "用户注册",
			"islogin":     islogin,
			"sessions":    sessions,
			"forums":      forums,
			"webname":     webname,
			"description": description,
			"forumname":   forumname,
		},
	)
}

func ConfirmEmail(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)
	// 网站描述
	webname := setting.ServerSetting.Sitename
	description := setting.ServerSetting.Sitebrief
	forumname := "邮箱验证"

	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"newemail_check.html",
		// Pass the data that the page uses
		gin.H{
			"title":       "验证邮箱",
			"islogin":     islogin,
			"sessions":    sessions,
			"forums":      forums,
			"webname":     webname,
			"description": description,
			"forumname":   forumname,
		},
	)
}

// https://127.0.0.1:9000/register/checkMail?email=wowiwo@yeah.net&time=1566011025&sign=3a6e2ad3aedf77a1dd7f8c3d9c532945
func CheckEmail(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)
	// 网站描述
	webname := setting.ServerSetting.Sitename
	description := setting.ServerSetting.Sitebrief
	forumname := "邮箱验证"

	email := c.DefaultQuery("email", "")
	time := c.DefaultQuery("time", "")
	sign := c.DefaultQuery("sign", "")

	// 链接地址错误
	if sign != util.EncodeMD5(email+time) {
		c.HTML(
			http.StatusOK,
			"email_check_err.html",
			gin.H{
				"title":       "验证邮箱",
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
			"email_check_pass_time.html",
			gin.H{
				"title":       "验证邮箱",
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
	user.UpdateEmailChecked(email)

	// 更新缓存
	session.SetSession(c, "emailchecked", "1")

	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"email_check_pass.html",
		// Pass the data that the page uses
		gin.H{
			"title":       "验证邮箱",
			"islogin":     islogin,
			"sessions":    sessions,
			"forums":      forums,
			"webname":     webname,
			"description": description,
			"forumname":   forumname,
		},
	)
}
