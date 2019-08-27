package web

import (
	"gorobbs/model"
	"gorobbs/service/v1/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//编辑主题
func EditPost(c *gin.Context) {
	postId, _ := strconv.Atoi(c.Param("id"))

	post, _ := model.GetPostById(postId)

	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)

	// 获取附件列表
	attachs, _ := model.GetAttachsByPostId(postId)

	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"editpost.html",
		// Pass the data that the page uses
		gin.H{
			"post":     post,
			"islogin":  islogin,
			"sessions": sessions,
			"forums":   forums,
			"attachs":  attachs,
		},
	)
}
