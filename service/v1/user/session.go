package user

import (
	"gorobbs/model"
	"gorobbs/package/session"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserSession struct {
	Username      string `json:"username"`
	Userid        int    `json:"userid"`
	Useravatar    string `json:"useravatar"`
	Useremail     string `json:"useremail"`
	EmailChecked  string `json:"email_checked"`
	Userpostcnt   int    `json:"userpostcnt"`
	Userthreadcnt int    `json:"userthreadcnt"`
	Isadmin       string `json:"isadmin"`
	WechatUnionID string `json:"wechat_union_id"`
}

// 登录设定必要的session信息
func LoginSession(c *gin.Context, user model.User, sok chan int) {
	emailChecked := strconv.Itoa(user.EmailChecked)
	session.SetSession(c, "username", user.Username)
	session.SetSession(c, "userid", strconv.Itoa(int(user.ID)))
	session.SetSession(c, "useravatar", user.Avatar)
	session.SetSession(c, "useremail", user.Email)
	session.SetSession(c, "emailchecked", emailChecked)
	session.SetSession(c, "userpostcnt", strconv.Itoa(user.PostsCnt))
	session.SetSession(c, "userthreadcnt", strconv.Itoa(user.ThreadsCnt))
	session.SetSession(c, "isadmin", IsAdmin(user.GroupID))
	session.SetSession(c, "wechat_union_id", user.WechatUnionID)
	sok <- 1
}

// 绑定微信时设定session
func SetWechatSession(c *gin.Context, wechatUnionID string) {
	session.SetSession(c, "wechat_union_id", wechatUnionID)
}

// 绑定微信时设定session
func GetWechatSession(c *gin.Context) {
	session.GetSession(c, "wechat_union_id")
}

// 获取session中信息
func GetSessions(c *gin.Context) (sessions *UserSession) {
	username := session.GetSession(c, "username")
	userid, _ := strconv.Atoi(session.GetSession(c, "userid"))
	useravatar := session.GetSession(c, "useravatar")
	useremail := session.GetSession(c, "useremail")
	emailchecked := session.GetSession(c, "emailchecked")
	userpostcnt, _ := strconv.Atoi(session.GetSession(c, "userpostcnt"))
	userthreadcnt, _ := strconv.Atoi(session.GetSession(c, "userthreadcnt"))
	isadmin := session.GetSession(c, "isadmin")
	wechatUnionID := session.GetSession(c, "wechat_union_id")

	sessions = &UserSession{
		Username:      username,
		Userid:        userid,
		Useravatar:    useravatar,
		Useremail:     useremail,
		EmailChecked:  emailchecked,
		Userpostcnt:   userpostcnt,
		Userthreadcnt: userthreadcnt,
		Isadmin:       isadmin,
		WechatUnionID: wechatUnionID,
	}

	return
}

// 登出操作session
func LogoutSession(c *gin.Context) {
	session.DeleteSession(c, "username")
	session.DeleteSession(c, "userid")
	session.DeleteSession(c, "useravatar")
	session.DeleteSession(c, "useremail")
	session.DeleteSession(c, "emailchecked")
	session.DeleteSession(c, "userpostcnt")
	session.DeleteSession(c, "userthreadcnt")
	session.DeleteSession(c, "isadmin")
	session.DeleteSession(c, "wechat_union_id")
}

// 判断是否已经登录
func IsLogin(c *gin.Context) (res bool) {
	username := session.GetSession(c, "username")

	if len(username) > 0 {
		res = true
	} else {
		res = false
	}

	return
}
