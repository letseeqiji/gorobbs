package web

import (
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

func Register(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)

	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"register.html",
		// Pass the data that the page uses
		gin.H{
			"title":    "用户注册",
			"islogin":  islogin,
			"sessions": sessions,
			"forums":   forums,
		},
	)
}
