package web

import (
	"gorobbs/model"
	"gorobbs/package/setting"
	"gorobbs/service/v1/forum"
	"gorobbs/service/v1/user"
	"gorobbs/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	tag_service "gorobbs/service/v1/tag"
)

// 分类模块页面
/*
查询指定fid模块下分页的帖子列表【按照更新日期倒叙排序】
查询全局置顶的帖子
查询当前模块指定的帖子
需要的条件：fid page
*/
func Forums(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)

	fid, _ := strconv.Atoi(c.Param("id"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	// 获取 thread 列表---需要优化
	threadList, _ := forum.GetThreadListByForumID(fid, page)
	threadTotle := forum.GetThreadTotleByForumID(fid)
	forumTopThreadList, _ := model.GetTopThreadsForum(fid)

	pages := util.Pagination("?page={page}", threadTotle, page, 20)
	forumInfo, _ := model.GetForumByID(fid)

	// 网站描述
	webname := setting.ServerSetting.Sitename
	description := setting.ServerSetting.Sitebrief
	forumname := forumInfo.Name

	//
	tags, _ := tag_service.GetTagForumsByForumID(fid)

	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"forumslist.html",
		// Pass the data that the page uses
		gin.H{
			"threadList":            threadList,
			"forum_top_thread_list": forumTopThreadList,
			"forums":                forums,
			"fid":                   fid,
			"islogin":               islogin,
			"sessions":              sessions,
			"pages":                 pages,
			"threadtotle":           threadTotle,
			"foruminfoxx":           forumInfo,
			"webname":               webname,
			"description":           description,
			"forumname":             forumname,
			"tags":                  tags,
		},
	)

}
