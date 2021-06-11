package v1

import (
	"gorobbs/model"
	"gorobbs/package/app"
	"gorobbs/package/file"
	"gorobbs/package/gredis"
	"gorobbs/package/logging"
	"gorobbs/package/rcode"
	"gorobbs/package/regex"
	"gorobbs/package/session"
	"gorobbs/package/upload"
	"gorobbs/package/validator"
	user_service "gorobbs/service/v1/user"
	"gorobbs/util"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/gommon/log"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

//获取用户 根据用户名
func GetUser(c *gin.Context) {
	err := gredis.Lpush("reg:email", util.GenRandCode(6)+"t@t.com", 5)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"msg":  err,
			"data": make(map[string]interface{}),
		})

		return
	}
	res, _ := gredis.Brpop("reg:email")

	if res == "" {
		time.Sleep(time.Second * 3)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1002,
		"msg":  "pop",
		"data": res,
	})

	return

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name := c.Query("name"); name != "" {
		maps["name"] = name
		data["user"], _ = model.GetUser(maps)
	}

	code := rcode.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  rcode.GetMessage(code),
		"data": data,
	})
}

/*
接受邮箱密码
验证邮箱密码都非空
根据邮箱查出用户信息
{
code:400
msg:邮箱或者密码错误
data:{}
}
验证密码的正确性
如果邮箱密码都对，则生成token返回客户端
{
code:200
msg:登录成功
data:{
	token:123123123
	exptime:23432
}
}
*/
func UserLogin(c *gin.Context) {
	// 1, 获取并验证参数
	email := c.DefaultPostForm("email", "")
	password := c.DefaultPostForm("password", "")
	code := rcode.INVALID_PARAMS
	isEmail := regex.IsEmail(email)

	valid := &validation.Validation{}
	if isEmail {
		user_service.LoginValidWithEmail(valid, email, password)
	} else {
		user_service.LoginValidWithName(valid, email, password)
	}
	if valid.HasErrors() {
		validator.VErrorMsg(c, valid, code)
		return
	}

	// 2，验证邮箱和密码
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if isEmail {
		maps["email"] = email
	} else {
		maps["username"] = email
	}

	user, err := model.GetUser(maps)
	if err != nil {
		code = rcode.ERROR_NOT_EXIST_USER
		app.JsonOkResponse(c, code, data)
		return
	}

	// 获取加密的密码
	hashPassword := user.Password
	if !util.VerifyString(password, hashPassword) {
		code = rcode.ERROR_NOT_EXIST_USER
		app.JsonOkResponse(c, code, data)
		return
	}

	// 3，验证通过 生成token和session
	code = rcode.SUCCESS

	// 生成session  使nginx报502错误
	var sok chan int = make(chan int, 1)
	go user_service.LoginSession(c, user, sok)
	<-sok

	app.JsonErrResponse(c, code)
}

// 退出
func UserLogout(c *gin.Context) {

	// 3，验证通过 生成token和session
	code := rcode.SUCCESS
	user_service.LogoutSession(c)

	app.JsonErrResponse(c, code)
}

//刷新token
func RefreshToken(c *gin.Context) {
	token := c.Query("token")

	newToken, time, _ := util.RefreshToken(token)
	data := make(map[string]interface{})
	data["token"] = newToken
	data["exp_time"] = time

	code := rcode.SUCCESS
	app.JsonOkResponse(c, code, data)
}

//新增用户
func AddUser(c *gin.Context) {
	username := c.PostForm("username")
	// 防止xss
	username = util.XssPolice(username)
	password := c.PostForm("password")
	// phone := c.PostForm("phone")
	email := c.PostForm("email")

	user := &model.User{}
	valid := &validation.Validation{}
	var err error

	code := rcode.INVALID_PARAMS

	user_service.AddUserValid(valid, username, password, email)
	if valid.HasErrors() {
		validator.VErrorMsg(c, valid, code)
		return
	}

	if !model.ExistUserByName(username) && !model.ExistUserByEmail(email) {
		code = rcode.SUCCESS
		ip := c.ClientIP()
		user, err = model.AddUser(username, password, email, ip)
		if err != nil {
			code = rcode.ERROR
			logging.Info("注册入库错误", err.Error())

			app.JsonErrResponse(c, code)
			return
		}
	} else {
		code = rcode.ERROR_EXIST_TAG
	}

	app.JsonOkResponse(c, code, user)
}

func AddUser2(c *gin.Context) {
	var user model.User
	code := rcode.SUCCESS

	// 只能绑定json传值  作为借口 可以   但是接受表单数据 不行
	if err := c.ShouldBind(&user); err != nil {
		code = rcode.ERROR_BIND_DATA
		app.JsonOkResponse(c, code, nil)
		return
	}
	user.Password, _ = util.BcryptString(user.Password)
	model.GetDb().Create(&user)

	app.JsonOkResponse(c, code, user)
}

//修改用户
func EditUser(c *gin.Context) {
}

//删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := rcode.SUCCESS
	// 验证管理员才可以
	uid, _ := strconv.Atoi(session.GetSession(c, "userid"))
	isadmin := user_service.IsAdmin(uid)
	if isadmin == "0" {
		code = rcode.UNPASS
		app.JsonErrResponse(c, code)
		return
	}

	// 删除用户前 要删除一些东西：用户的帖子，用户的评论，用户的收藏，凡是有用户id 的表内容都要删除
	// 还是说如果用户下面有很多内容就不让删除，只能删除新建的用户
	err := user_service.DelUserByID(id)
	if err != nil {
		log.Print("api.v1.user.deluser.deluserbyid:err:", code)

		code = rcode.ERROR_SQL_DELETE_FAIL
		app.JsonErrResponse(c, code)
		return
	}

	app.JsonOkResponse(c, code, nil)
}

func ResetUserPassword(c *gin.Context) {
	oldpassword := c.PostForm("password_old")
	newpassword := c.PostForm("password_new")
	uid := c.Param("id")
	code := rcode.SUCCESS

	// 验证原来的密码正确性
	maps := make(map[string]interface{})

	maps["id"] = uid
	user, err := model.GetUser(maps)
	if err != nil {
		code = rcode.ERROR_NOT_EXIST_USER
		app.JsonErrResponse(c, code)
		return
	}

	// 获取加密的密码
	hashPassword := user.Password
	if !util.VerifyString(oldpassword, hashPassword) {
		code = rcode.ERROR
		app.JsonErrResponse(c, code)
		return
	}

	user.Password, _ = util.BcryptString(newpassword)
	err = user_service.ResetPassword(user.Password, int(user.ID))
	if err != nil {
		code = rcode.ERROR
		app.JsonErrResponse(c, code)
		return
	}

	app.JsonOkResponse(c, code, nil)
}

func ResetUserAvatar(c *gin.Context) {
	userAvatar, err := c.FormFile("avatar")
	uid, err := strconv.Atoi(c.Param("id"))
	fileName := userAvatar.Filename
	code := rcode.SUCCESS

	if err != nil {
		code = rcode.ERROR
		app.JsonErrResponse(c, code)
		return
	}
	// 限制图片的格式 和 大小
	if !upload.CheckImageExt(fileName) {
		code = rcode.ERROR_IMAGE_BAD_EXT
		app.JsonErrResponse(c, code)
		return
	}

	if !upload.CheckImageSize2(userAvatar) {
		code = rcode.ERROR_IMAGE_TOO_LARGE
		app.JsonErrResponse(c, code)
		return
	}

	filePath := "upload/avatar/" + c.Param("id")
	// 判断路径是否存在 不存在则创建
	filePath, err = file.CreatePathInToday(filePath)
	if err != nil {
		code = rcode.ERROR_FILE_CREATE_FAIL
		app.JsonErrResponse(c, code)
		return
	}

	fullFileName := filePath + "/" + fileName
	err = c.SaveUploadedFile(userAvatar, fullFileName)
	if err != nil {
		code = rcode.ERROR_FILE_SAVE_FAIL
		app.JsonErrResponse(c, code)
		return
	}

	err = user_service.ResetAvatar(fullFileName, uid)
	if err != nil {
		code = rcode.ERROR
		app.JsonErrResponse(c, code)
		return
	}

	// 更新session
	session.SetSession(c, "useravatar", "/"+fullFileName)

	app.JsonOkResponse(c, code, fullFileName)
}

func ResetUserName(c *gin.Context) {
	userName := c.PostForm("user_name")
	uid, _ := strconv.Atoi(c.Param("id"))
	code := rcode.SUCCESS

	err := user_service.ResetName(userName, uid)
	if err != nil {
		code = rcode.ERROR_SQL_UPDATE_FAIL
		app.JsonErrResponse(c, code)
		return
	}

	// 更新session
	session.SetSession(c, "username", userName)

	app.JsonOkResponse(c, code, nil)
}

// 检测用户名是否被使用
func CheckNameUsed(c *gin.Context) {
	name := c.DefaultQuery("username", "")
	code := rcode.SUCCESS
	data := make(map[string]interface{})

	if model.ExistUserByName(name) {
		data["is_used"] = 1
	} else {
		data["is_used"] = 0
	}

	app.JsonOkResponse(c, code, data)
}

// 检测邮箱是否被使用
func CheckEmailUsed(c *gin.Context) {
	email := c.DefaultQuery("email", "")
	code := rcode.SUCCESS
	data := make(map[string]interface{})

	if model.ExistUserByEmail(email) {
		data["is_used"] = 1
	} else {
		data["is_used"] = 0
	}

	app.JsonOkResponse(c, code, data)
}

// 检测邮箱是否被使用
func CheckPhoneUsed(c *gin.Context) {
	phone := c.DefaultQuery("phone", "")
	code := rcode.SUCCESS
	data := make(map[string]interface{})

	if model.ExistUserByPhone(phone) {
		data["is_used"] = 1
	} else {
		data["is_used"] = 0
	}

	app.JsonOkResponse(c, code, data)
}

func IsEmailChecked(c *gin.Context) {
	email := c.DefaultPostForm("email", "")
	code := rcode.SUCCESS
	if len(email) == 0 {
		code = rcode.ERROR_UNFIND_DATA
		app.JsonErrResponse(c, code)
		return
	}

	user, err := model.GetUser(map[string]interface{}{"email": email})
	if err != nil {
		code = rcode.ERROR
		app.JsonErrResponse(c, code)
		return
	}

	app.JsonOkResponse(c, code, user.EmailChecked)
}
