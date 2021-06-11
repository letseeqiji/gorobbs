package web

import (
	"gorobbs/package/setting"
	"gorobbs/service/v1/user"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

/*var (
	forums []model.Forum
)*/

/*func init() {
	forums = layout_service.GetForumList()
}*/

func Login(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)
	// 网站描述
	webname := setting.ServerSetting.Sitename
	description := setting.ServerSetting.Sitebrief
	forumname := "登录"

	// 第三方登录
	// 微信
	wechatAppID := setting.WechatSetting.AppID
	wechatCallBack := setting.WechatSetting.CallBackURL
	wechatCallBack = url.QueryEscape(wechatCallBack)

	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"login.html",
		// Pass the data that the page uses
		gin.H{
			"title":          "Home Page",
			"islogin":        islogin,
			"sessions":       sessions,
			"forums":         forums,
			"description":    description,
			"forumname":      forumname,
			"webname":        webname,
			"wechatAppID":    wechatAppID,
			"wechatCallBack": wechatCallBack,
		},
	)
}
