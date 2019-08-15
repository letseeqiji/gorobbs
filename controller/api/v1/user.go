package v1

import (
	"gorobbs/model"
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
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  rcode.GetMessage(code),
			"data": data,
		})
		return
	}

	// 获取加密的密码
	hashPassword := user.Password
	if !util.VerifyString(password, hashPassword) {
		code = rcode.ERROR_NOT_EXIST_USER
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  rcode.GetMessage(code),
			"data": data,
		})
		return
	}

	/*c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  rcode.GetMessage(code),
		"data": bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(user.Password)),
	})
	return*/

	// 3，验证通过 生成token和session
	code = rcode.SUCCESS

	// 生成session  使nginx报502错误
	var sok chan int = make(chan int, 1)
	go user_service.LoginSession(c, user, sok)
	<- sok

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": rcode.GetMessage(code),
		//"data":    data,
	})

}

// 退出
func UserLogout(c *gin.Context) {

	// 3，验证通过 生成token和session
	code := rcode.SUCCESS
	user_service.LogoutSession(c)

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": rcode.GetMessage(code),
		"data":    make(map[string]interface{}),
	})
}

//刷新token
func RefreshToken(c *gin.Context) {
	token := c.Query("token")

	newToken, time, _ := util.RefreshToken(token)
	data := make(map[string]interface{})
	data["token"] = newToken
	data["exp_time"] = time

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  rcode.GetMessage(200),
		"data": data,
	})
}

//新增用户
func AddUser(c *gin.Context) {
	username := c.PostForm("username")
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
			logging.Info("注册入库错误",err.Error())

			c.JSON(http.StatusInternalServerError, gin.H{
				"code": code,
				"msg":  rcode.GetMessage(code),
			})
			return
		}
	} else {
		code = rcode.ERROR_EXIST_TAG
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  rcode.GetMessage(code),
		"data": user,
	})
}

func AddUser2(c *gin.Context) {
	var user model.User
	// 只能绑定json传值  作为借口 可以   但是接受表单数据 不行
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(
			500,
			"绑定错误",
		)
	}
	user.Password, _ = util.BcryptString(user.Password)
	model.GetDb().Create(&user)

	code := rcode.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  rcode.GetMessage(code),
		"data": &user,
	})
}

//修改用户
func EditUser(c *gin.Context) {
}

//删除用户
func DeleteUser(c *gin.Context) {
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
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": rcode.GetMessage(code),
		})
		return
	}

	// 获取加密的密码
	hashPassword := user.Password
	if !util.VerifyString(oldpassword, hashPassword) {
		code = rcode.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": rcode.GetMessage(code),
		})
		return
	}

	user.Password, _ = util.BcryptString(newpassword)
	err = user_service.ResetPassword(user.Password, int(user.ID))
	if err != nil {
		code = rcode.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": rcode.GetMessage(code),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": rcode.GetMessage(code),
	})
}

/*func ResetUserAvatar(c *gin.Context) {
	code := e.SUCCESS
	data := ""

	imgfile, image, err := c.Request.FormFile("avatar")
	if err != nil {
		logging.Warn(err)
		code = e.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
	}

	if image == nil {
		code = e.INVALID_PARAMS
	} else {
		imageName := upload.GetImageName(image.Filename)
		fullPath := "upload/avatar/" + c.Param("id")
		// 判断路径是否存在 不存在则创建
		err = file.CreatePath(fullPath)
		if err != nil {
			c.JSON(200, gin.H{
				"code":    500,
				"message": err.Error(),
			})
			return
		}

		src := fullPath + "/" + imageName
		if ! upload.CheckImageExt(imageName) || ! upload.CheckImageSize(imgfile) {
			code = e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT
		} else {
			err := upload.CheckImage(fullPath)
			if err != nil {
				logging.Warn(err)
				code = e.ERROR_UPLOAD_CHECK_IMAGE_FAIL
			} else if err := c.SaveUploadedFile(image, src); err != nil {
				logging.Warn(err)
				code = e.ERROR_UPLOAD_SAVE_IMAGE_FAIL
			} else {
				data = src
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}*/

func ResetUserAvatar(c *gin.Context) {
	userAvatar, err := c.FormFile("avatar")
	uid, _ := strconv.Atoi(c.Param("id"))
	fileName := userAvatar.Filename
	//code := rcode.SUCCESS

	if err != nil {
		c.JSON(200, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}
	// 限制图片的格式 和 大小
	if ! upload.CheckImageExt(fileName)  {
		c.JSON(200, gin.H{
			"code":    403,
			"message": "图片格式不正确",
		})
		return
	}

	if !upload.CheckImageSize2(userAvatar) {
		c.JSON(200, gin.H{
			"code":    403,
			"message": "图片大小超标了",
		})
		return
	}

	filePath := "upload/avatar/" + c.Param("id")
	// 判断路径是否存在 不存在则创建
	filePath, err = file.CreatePathInToday(filePath)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	fullFileName := filePath + "/" + fileName
	err = c.SaveUploadedFile(userAvatar, fullFileName)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	err = user_service.ResetAvatar(fullFileName, uid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": err.Error(),
		})
		return
	}

	// 更新session
	session.SetSession(c, "useravatar", "/"+fullFileName)

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "修改成功",
		"data":    fullFileName,
		"filesize":userAvatar.Size,
		//1/Video_2019-07-25_102727.wmv.png", filesize: 26650725, message: "修改成功" }
	})
}

func ResetUserName(c *gin.Context) {
	userName := c.PostForm("user_name")
	uid, _ := strconv.Atoi(c.Param("id"))

	err := user_service.ResetName(userName, uid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": err.Error(),
		})
		return
	}

	// 更新session
	session.SetSession(c, "username", userName)

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "修改成功",
	})
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

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  rcode.GetMessage(code),
		"data": data,
	})
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

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  rcode.GetMessage(code),
		"data": data,
	})
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

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  rcode.GetMessage(code),
		"data": data,
	})
}
