package web

import (
	"gorobbs/model"
	"gorobbs/package/setting"
	thread_service "gorobbs/service/v1/thread"
	"gorobbs/service/v1/user"
	"html"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*func init()  {
	forums = layout_service.GetForumList()
}*/

// 帖子详情页
func Thread(c *gin.Context) {

	threadId, _ := strconv.Atoi(c.Param("id"))
	// c.JSON(200, gin.H{"id": threadId})

	thread, _ := model.GetThreadById(threadId)
	fpost, _ := model.GetThreadFirstPostByTid(threadId)

	// c.JSON(200, gin.H{"data": html.UnescapeString(fpost.Message), "yd": fpost.Message})

	fpost.MessageFmt = html.UnescapeString(fpost.MessageFmt)

	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)

	// 获取平路列表
	postlist, _ := model.GetThreadPostListByTid(threadId, 500, 1)
	postlistLen := len(postlist)

	// 获取附件列表
	attachs, _ := model.GetAttachsByPostId(int(fpost.ID))

	// 获取是否收藏
	isfav, _ := model.CheckFavourite(sessions.Userid, threadId)
	//isfav:= 1
	//isLiked, _ := model.CheckPostlike(sessions.Userid, )

	// 获取用户的最新的threads
	userNewestThreads, _ := thread_service.GetUserThreads(thread.UserID)

	// 阅读量增加1
	model.UpdateThreadViewsCnt(threadId)

	// 网站描述
	description := setting.ServerSetting.Sitebrief
	forumname := thread.Subject

	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"thread.html",
		// Pass the data that the page uses
		gin.H{
			"thread":  thread,
			"fpost":   fpost,
			"islogin": islogin,
			"sessions": sessions,
			"postlist":postlist,
			"post_list_len":postlistLen,
			"forums":forums,
			"user_newest_threads":userNewestThreads,
			"attachs":attachs,
			"isfav":isfav,
			"description":description,
			"forumname":forumname,
		},
	)
}

// ThreadAddPost 高级回复也
func ThreadAddPost(c *gin.Context)  {
	threadId, _ := strconv.Atoi(c.Param("id"))
	sessions := user.GetSessions(c)
	islogin := user.IsLogin(c)
	// 网站描述
	description := setting.ServerSetting.Sitebrief
	forumname := "高级回复"
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"advance_post.html",
		// Pass the data that the page uses
		gin.H{
			"sessions": sessions,
			"islogin": islogin,
			"forums":forums,
			"thread_id":  threadId,
			"description":description,
			"forumname":forumname,
		},
	)
}

// 发新主题
func NewThread(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)

	// 网站描述
	description := setting.ServerSetting.Sitebrief
	forumname := "新主题"

	// 获取平路列表
	//forums, _ := model.GetForumsList("id asc")

	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"newthread.html",
		// Pass the data that the page uses
		gin.H{
			"forums":  forums,
			"islogin": islogin,
			"sessions": sessions,
			"description":description,
			"forumname":forumname,
		},
	)
}

//编辑主题
func EditThread(c *gin.Context) {
	threadId, _ := strconv.Atoi(c.Param("id"))
	thread, _ := model.GetThreadById(threadId)
	fpost, _ := model.GetThreadFirstPostByTid(threadId)

	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)

	// 获取附件列表
	attachs, _ := model.GetAttachsByPostId(int(fpost.ID))

	// 网站描述
	webname := setting.ServerSetting.Sitename
	description := setting.ServerSetting.Sitebrief
	forumname := thread.Subject

	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"editthread.html",
		// Pass the data that the page uses
		gin.H{
			"thread":  thread,
			"fpost":   fpost,
			"islogin": islogin,
			"sessions": sessions,
			"forums":forums,
			"attachs":attachs,
			"webname":webname,
			"description":description,
			"forumname":forumname,
		},
	)
}




