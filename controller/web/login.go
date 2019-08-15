package web

import (
	"gorobbs/package/setting"
	"gorobbs/service/v1/user"
	"net/http"

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
	/*if islogin {
		c.Redirect(http.StatusMovedPermanently, "/")
		return
		//c.JSON(200, gin.H{"code":403, "msg":"login"})
	}*/ /*else {
		c.JSON(200, gin.H{"code":200, "msg":"no login"})
	}
	return*/
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"login.html",
		// Pass the data that the page uses
		gin.H{
			"title":    "Home Page",
			"islogin":  islogin,
			"sessions": sessions,
			"forums":   forums,
			"description":description,
			"forumname":forumname,
			"webname":webname,
		},
	)
}
