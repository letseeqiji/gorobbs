package admin

import (
	"gorobbs/model"
	"gorobbs/package/app"
	"gorobbs/package/rcode"
	"gorobbs/service/v1/user"
	"gorobbs/util"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func AdminUserList(c *gin.Context) {
	users, _ := model.GetUsers(50, "created_at desc", map[string]interface{}{})
	sessions := user.GetSessions(c)
	c.HTML(
		200,
		"user_list.html",
		gin.H{
			"users":    users,
			"sessions": sessions,
		})
}

func AdminGroupList(c *gin.Context) {
	sessions := user.GetSessions(c)
	groupLis, _ := model.GetUserGroupList()
	c.HTML(200,
		"group_list.html",
		gin.H{
			"groupList": groupLis,
			"sessions":  sessions,
		})
}

func AdminUserCreate(c *gin.Context) {
	sessions := user.GetSessions(c)
	groupLis, _ := model.GetUserGroupList()
	c.HTML(200,
		"user_add.html",
		gin.H{
			"groupList": groupLis,
			"sessions":  sessions,
		})
}

// 需要字段
// Username
//GroupID
//Email
//Password
// email=sdsdsdsdsdsdsdsdsdsdsdsdsd&username=wowiwo6%40yeah.net&password=123456&group_id=0
func AdminUserAdd(c *gin.Context) {
	email := c.PostForm("email")
	username := c.PostForm("username")
	password := c.PostForm("password")
	group_id, _ := strconv.Atoi(c.PostForm("group_id"))
	ip := c.ClientIP()

	_, err := model.AddUserPro(&model.User{
		Username:  username,
		GroupID:   group_id,
		Email:     email,
		Password:  password,
		CreateIp:  ip,
		LoginDate: time.Now(),
	})
	if err != nil {
		c.JSON(200, gin.H{
			"code":    500,
			"message": "保存失败",
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "保存成功",
	})
}

func AdminUserEdit(c *gin.Context) {
	sessions := user.GetSessions(c)
	uid, _ := strconv.Atoi(c.Query("id"))
	groupLis, _ := model.GetUserGroupList()
	userInfo, _ := model.GetUserByID(uid)
	c.HTML(200,
		"user_edit.html",
		gin.H{
			"userinfo":  userInfo,
			"groupList": groupLis,
			"sessions":  sessions,
		})
}

func AdminUserUpdate(c *gin.Context) {
	uid := c.PostForm("userid")
	email := c.PostForm("email")
	username := c.PostForm("username")
	password := c.PostForm("password")
	group_id, _ := strconv.Atoi(c.PostForm("group_id"))
	code := rcode.SUCCESS

	var wmap = make(map[string]interface{})
	var updateItems = make(map[string]interface{})
	wmap["id"] = uid
	updateItems["email"] = email
	updateItems["username"] = username
	updateItems["group_id"] = group_id
	if len(password) != 0 {
		password, _ = util.BcryptString(password)
		updateItems["password"] = password
	}

	err := model.UpdateUser(wmap, updateItems)

	if err != nil {
		code = rcode.ERROR_SQL_UPDATE_FAIL
		app.JsonErrResponse(c, code)
		return
	}

	app.JsonOkResponse(c, code, nil)
}
