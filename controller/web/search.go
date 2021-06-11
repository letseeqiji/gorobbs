package web

import (
	"gorobbs/model"
	"gorobbs/package/setting"
	"gorobbs/service/v1/user"
	searchtool "gorobbs/tools/search"
	"net/http"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	// 用户是否登录
	islogin := user.IsLogin(c)
	// 获取sessions
	sessions := user.GetSessions(c)
	// 网站描述
	webname := setting.ServerSetting.Sitename
	description := setting.ServerSetting.Sitebrief
	forumname := "搜索"

	var threads []*model.Thread

	key := c.DefaultQuery("key", "")

	if utf8.RuneCountInString(key) > 1 {
		output := searchtool.Search(key, 1)
		// 搜索输出格式见types.SearchResponse结构体
		ids := searchtool.OutPutIds(output)
		threads, _ = model.GetThreadsByIDs(ids)
	}

	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"search.html",
		// Pass the data that the page uses
		gin.H{
			"forums":      forums,
			"islogin":     islogin,
			"sessions":    sessions,
			"webname":     webname,
			"description": description,
			"forumname":   forumname,
			"keyword":     key,
			"threads":     threads,
		},
	)
}