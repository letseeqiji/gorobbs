package web

import (
	"gorobbs/model"
	package_online "gorobbs/package/online"
	"gorobbs/package/setting"
	"gorobbs/service/v1/user"
	"gorobbs/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	layout_service "gorobbs/service/v1/layout"
)

var (
	forums []model.Forum
)

const PAGE_SIZE int = 20

func init() {
	forums = layout_service.GetForumList()
}

func Index(c *gin.Context) {
	// 获取分页
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	// 根据分页获取帖子列表
	threadList, _ := model.GetThreadList(page)
	// 获取置顶的帖子列表
	topThreadList, _ := model.GetTopThreadsWholeWebSite()
	// 获取帖子总数
	threadTotle, _ := model.GetThreadTotleCount()
	// 用户是否登录
	islogin := user.IsLogin(c)
	// 获取sessions
	sessions := user.GetSessions(c)
	// 制作分页效果
	pages := util.Pagination("?page={page}", threadTotle, page, PAGE_SIZE)
	// 获取最新会员
	newestUser, _ := user.GetNewestTop12Users()

	// 获取在线人数
	online, _ := package_online.DbSize()
	threadsNum, _ := model.CountThreadsNum()
	//帖子数：8
	postsNum, _ := model.CountPostNum()
	//用户数：2
	usersNum, _ := model.CountUserNum()
	// 网站描述
	webname := setting.ServerSetting.Sitename
	description := setting.ServerSetting.Sitebrief
	forumname := "首页"
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses
		gin.H{
			"threadList":      threadList,
			"top_thread_list": topThreadList,
			"forums":          forums,
			"islogin":         islogin,
			"sessions":        sessions,
			"newestuser":      newestUser,
			"pages":           pages,
			"online":          online,
			"threads_num":     threadsNum,
			"posts_num":       postsNum,
			"users_num":       usersNum,
			"webname":         webname,
			"description":     description,
			"forumname":       forumname,
		},
	)
}
