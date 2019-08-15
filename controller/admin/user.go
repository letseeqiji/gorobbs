package admin

import (
	"github.com/gin-gonic/gin"
	"gorobbs/model"
	"gorobbs/service/v1/user"
	"strconv"
	"time"
)

func AdminUserList(c *gin.Context)  {
	users, _ := model.GetUsers(50, "created_at desc", map[string]interface{}{})
	sessions := user.GetSessions(c)
	c.HTML(
		200,
		"user_list.html",
		gin.H{
			"users":users,
			"sessions":sessions,
		})
}

func AdminGroupList(c *gin.Context)  {
	sessions := user.GetSessions(c)
	groupLis, _ := model.GetUserGroupList()
	c.HTML(200,
		"group_list.html",
		gin.H{
			"groupList":groupLis,
			"sessions":sessions,
		})
}

func AdminUserCreate(c *gin.Context)  {
	sessions := user.GetSessions(c)
	groupLis, _ := model.GetUserGroupList()
	c.HTML(200,
		"user_add.html",
		gin.H{
			"groupList":groupLis,
			"sessions":sessions,
	})
}

// 需要字段
// Username
//GroupID
//Email
//Password
// email=sdsdsdsdsdsdsdsdsdsdsdsdsd&username=wowiwo6%40yeah.net&password=123456&group_id=0
func AdminUserAdd(c *gin.Context)  {
	email := c.PostForm("email")
	username := c.PostForm("username")
	password := c.PostForm("password")
	group_id,_ := strconv.Atoi(c.PostForm("group_id"))
	ip := c.ClientIP()

	_, err := model.AddUserPro(&model.User{
		Username:username,
		GroupID:group_id,
		Email:email,
		Password:password,
		CreateIp:ip,
		LoginDate:time.Now(),
	})
	if err != nil {
		c.JSON(200, gin.H{
			"code":    500,
			"message": "保存失败",
		})
		return
	}

	c.JSON(200, gin.H{
		"code":200,
		"message":"保存成功",
		/*"message":map[string]interface{}{
			"sitename" :sitename,
			"sitebrief" :sitebrief,
			"runlevel" :runlevel,
			"user_create_on" :user_create_on,
			"user_create_email_on":user_create_email_on,
			"user_resetpw_on" :user_resetpw_on,
			"lang" :lang,
		},*/
	})
}


