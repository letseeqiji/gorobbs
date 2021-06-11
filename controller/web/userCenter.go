package web

import (
	"gorobbs/model"
	"gorobbs/package/setting"
	"gorobbs/service/v1/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	MY_INFO = iota
	MY_PASSWORD
	MY_AVATAR
	MY_NAME
	MY_EMAIL
	MY_THREADS
	MY_FAVS
	MY_DIGESTS
	MY_POSTS
	MY_NOTICE
)

const (
	U_INFO = iota
	U_THREADS
	U_POSTS
)

var (
	webname     string
	description string
	forumname   string
)

func init() {
	// 网站描述
	webname = setting.ServerSetting.Sitename
	description = setting.ServerSetting.Sitebrief
	forumname = "用户中心"
}

func MyInfo(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)

	uid := sessions.Userid
	userinfo, _ := user.GetUserByID(uid)

	tpl := "my_info.html"
	c.HTML(200, tpl, gin.H{
		"islogin":     islogin,
		"sessions":    sessions,
		"userinfo":    userinfo,
		"forums":      forums,
		"webname":     webname,
		"description": description,
		"forumname":   forumname,
	})
}

func MyPassword(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)

	tpl := "my_password.html"
	c.HTML(200, tpl, gin.H{
		"islogin":     islogin,
		"sessions":    sessions,
		"forums":      forums,
		"webname":     webname,
		"description": description,
		"forumname":   forumname,
	})
}

func MyAvatar(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)

	tpl := "my_avatar.html"
	c.HTML(200, tpl, gin.H{
		"islogin":     islogin,
		"sessions":    sessions,
		"forums":      forums,
		"webname":     webname,
		"description": description,
		"forumname":   forumname,
	})
}

func MyUsername(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)

	tpl := "my_name.html"
	c.HTML(200, tpl, gin.H{
		"islogin":     islogin,
		"sessions":    sessions,
		"forums":      forums,
		"webname":     webname,
		"description": description,
		"forumname":   forumname,
	})
}

func MyThread(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	uid := sessions.Userid

	// 找出uid下面的所有的thread--按日期倒序  用mythread模型即可
	//model.Mythread.GetMyThreads(uid int)
	myThreads, _ := model.GetMyThreadList(uid, page, 20, "created_at desc")

	tpl := "my_thread.html"
	c.HTML(200, tpl, gin.H{
		"islogin":     islogin,
		"sessions":    sessions,
		"forums":      forums,
		"mythreads":   myThreads,
		"webname":     webname,
		"description": description,
		"forumname":   forumname,
	})
}

func MyFavorite(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)
	uid := sessions.Userid
	favThreads, _ := model.GetMyFavouriteList(uid, 1, 200, "created_at desc")

	tpl := "my_favorite.html"
	c.HTML(200, tpl, gin.H{
		"islogin":     islogin,
		"sessions":    sessions,
		"forums":      forums,
		"threads":     favThreads,
		"webname":     webname,
		"description": description,
		"forumname":   forumname,
	})
}

func MyDigest(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)

	tpl := "my_digest.html"
	c.HTML(200, tpl, gin.H{
		"islogin":     islogin,
		"sessions":    sessions,
		"forums":      forums,
		"webname":     webname,
		"description": description,
		"forumname":   forumname,
	})
}

func MyPost(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	uid := sessions.Userid

	posts, _ := model.GetMyPostList(uid, page, 20, "created_at desc")

	tpl := "my_post.html"
	c.HTML(200, tpl, gin.H{
		"islogin":     islogin,
		"sessions":    sessions,
		"forums":      forums,
		"posts":       posts,
		"webname":     webname,
		"description": description,
		"forumname":   forumname,
	})
}

func UserInfo2(c *gin.Context) {
	action, _ := strconv.Atoi(c.DefaultQuery("action", "0"))

	var tpl string

	switch action {
	case U_THREADS:
		tpl = "u_thread.html" //?page=
	case U_POSTS:
		tpl = "u_posts.html"
	default:
		tpl = "u_info.html"
	}

	c.HTML(200, tpl, gin.H{})
}

func UserInfo(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)

	uid, _ := strconv.Atoi(c.Param("id"))
	userinfo, _ := user.GetUserByID(uid)

	tpl := "u_info.html"
	c.HTML(200, tpl, gin.H{
		"islogin":     islogin,
		"sessions":    sessions,
		"userinfo":    userinfo,
		"forums":      forums,
		"webname":     webname,
		"description": description,
		"forumname":   forumname,
	})
}

func UserThread(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	uid, _ := strconv.Atoi(c.Param("id"))
	userinfo, _ := user.GetUserByID(uid)

	// 找出uid下面的所有的thread--按日期倒序  用mythread模型即可
	//model.Mythread.GetMyThreads(uid int)
	myThreads, _ := model.GetMyThreadList(uid, page, 20, "created_at desc")

	tpl := "u_threads.html"
	c.HTML(200, tpl, gin.H{
		"islogin":     islogin,
		"sessions":    sessions,
		"forums":      forums,
		"mythreads":   myThreads,
		"userinfo":    userinfo,
		"webname":     webname,
		"description": description,
		"forumname":   forumname,
	})
}

func UserPost(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	uid, _ := strconv.Atoi(c.Param("id"))
	userinfo, _ := user.GetUserByID(uid)

	posts, _ := model.GetMyPostList(uid, page, 20, "created_at desc")

	tpl := "u_posts.html"
	c.HTML(200, tpl, gin.H{
		"islogin":     islogin,
		"sessions":    sessions,
		"forums":      forums,
		"posts":       posts,
		"userinfo":    userinfo,
		"webname":     webname,
		"description": description,
		"forumname":   forumname,
	})
}
