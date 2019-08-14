package v1

import (
	"github.com/aviddiviner/gin-limit"
	"github.com/gin-gonic/gin"
	"github.com/tommy351/gin-sessions"
	adminservice "gorobbs/controller/admin"
	apiservice "gorobbs/controller/api/v1"
	webservice "gorobbs/controller/web"
	"gorobbs/middleware/cros"
	"gorobbs/middleware/jwt"
	"gorobbs/middleware/loger"
	"gorobbs/middleware/online"
	"gorobbs/middleware/xss"
	"gorobbs/model"
	package_redis "gorobbs/package/gredis"
	"gorobbs/package/setting"
	"html/template"
	"net/http"
)

func InitRouter() *gin.Engine {
	// 禁用控制台颜色，当你将日志写入到文件的时候，你不需要控制台颜色。
	//gin.DisableConsoleColor()

	// 写入日志的文件
	/*f, _ := os.Create("log/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)*/

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cros.Cors())
	r.Use(limit.MaxAllowed(100))
	r.Use(online.OnLine())

	// 引入session
	store := sessions.NewCookieStore([]byte("secret123"))
	r.Use(sessions.Middleware("my_session", store))

	r.Use(loger.LoggerToFile())

	gin.SetMode(setting.ServerSetting.RunMode)

	// 模板函数
	r.SetFuncMap(template.FuncMap{
		"unescaped":   unescaped,
		"strtime":     StrTime,
		"plus1":       selfPlus,
		"numplusplus": numPlusPlus,
		"strip":		Long2IPString,
	})

	// 避免404
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", gin.H{})
	})
	r.NoMethod(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", gin.H{})
	})

	r.LoadHTMLGlob("views/*/**/***")
	// 推荐使用绝对路径 相当于简历了软连接--快捷方式
	r.StaticFS("/static", http.Dir("./static"))
	r.StaticFS("/upload", http.Dir("./upload"))

	// 用户前端页面
	web := r.Group("")
	{
		// 注册页
		web.GET("/register.html", webservice.Register)
		// 登录页
		web.GET("/login.html", webservice.Login)
		// 登出页面
		web.GET("/logout", apiservice.UserLogout)

		// 首页
		web.GET("/", webservice.Index)
		web.GET("/index.html", webservice.Index)
		web.GET("/default.html", webservice.Index)

		// 模块：主题列表页
		web.GET("/forum/:id/list.html", webservice.Forums)
		// 主题：新建页
		web.GET("/newthread.html", webservice.NewThread)
		// 主题：编辑页
		web.GET("/thread/:id/edit.html", webservice.EditThread)
		// 主题：详情页
		web.GET("/thread/:id/detail.html", webservice.Thread)
		// 高级回复
		web.GET("/thread/:id/apost.html", webservice.ThreadAddPost)
		web.GET("/post/:id/edit.html", webservice.EditPost)

		// 用户中心相关页面
		// 我的信息概览
		web.GET("/my.html", webservice.MyInfo)
		// 修改密码
		web.GET("/my_password.html", webservice.MyPassword)
		// 修改头像
		web.GET("/my_avatar.html", webservice.MyAvatar)
		// 修改用户名
		web.GET("/my_rename.html", webservice.MyUsername)
		// 我的帖子列表
		web.GET("/my_thread.html", webservice.MyThread)
		// 我的收藏列表
		web.GET("/my_favorite.html", webservice.MyFavorite)
		// 我的精华列表
		web.GET("/my_digest.html", webservice.MyDigest)
		// 我的回帖列表
		web.GET("/my_post.html", webservice.MyPost)

		// 查看其它用户内容
		web.GET("/user/:id/info.html", webservice.UserInfo)
		web.GET("/user/:id/thread.html", webservice.UserThread)
		web.GET("/user/:id/post.html", webservice.UserPost)

		// 前台管理员对帖子进行操作的模态框
		web.GET("/mod/thread/move.html", webservice.MoveMod)
		web.GET("/mod/thread/top.html", webservice.TopMod)
		web.GET("/mod/thread/delete.html", webservice.DeleteMod)
		web.GET("/mod/thread/close.html", webservice.CloseMod)
	}

	// 数据操作的接口
	apiv1 := r.Group("/api/v1")
	{
		// 检测用户名是否被使用
		apiv1.GET("/checkname", apiservice.CheckNameUsed)
		// 检测邮箱是否被使用
		apiv1.GET("/checkemail", apiservice.CheckEmailUsed)
		// 获取某用户
		apiv1.GET("/user", apiservice.GetUser)
		//注册
		apiv1.POST("/user", apiservice.AddUser)
		apiv1.POST("/register", apiservice.AddUser)
		// 登录
		apiv1.POST("/login", apiservice.UserLogin)
		// 登出操作
		apiv1.GET("/logout", apiservice.UserLogout)
		// 刷新token
		apiv1.GET("/token", apiservice.RefreshToken)
		// 更新用户
		apiv1.PUT("/user/:id", apiservice.EditUser)
		// 删除用户
		apiv1.DELETE("/user/:id", apiservice.DeleteUser)
		// 用户：重设密码
		apiv1.POST("/user/:id/password/reset", apiservice.ResetUserPassword)
		// 用户：重设头像
		apiv1.POST("/user/:id/avatar/reset", apiservice.ResetUserAvatar)
		// 用户：重设用户名
		apiv1.POST("/user/:id/name/reset", apiservice.ResetUserName)
		// 主题：发表
		apiv1.POST("/thread", xss.XSS(), apiservice.AddThread)
		// 主题：发表
		apiv1.POST("/thread/:id/favourite", apiservice.Addthreadfavourite)
		// 主题：删除
		apiv1.POST("/thread/:id/delete", apiservice.DeleteThreads)
		// 主题：移动到置顶模块
		apiv1.POST("/thread/:id/move", apiservice.MoveThreads)
		// 主题：置顶
		apiv1.POST("/thread/:id/top", apiservice.TopThreads)
		// 主题：关闭
		apiv1.POST("/thread/:id/close", apiservice.CloseThreads)
		// 主题：修改
		apiv1.POST("/thread/:id/update", apiservice.UpdateThread)
		// 添加评论
		apiv1.POST("/thread/:id/post", apiservice.AddPost)
		// 添加附件
		apiv1.POST("/thread/:id/attach/add", apiservice.AddthreadAttach)
		// 删除附件
		apiv1.POST("/thread/:id/attach/del", apiservice.DelthreadAttach)
		// 评论的相关操作
		apiv1.POST("/post/:id/update", apiservice.UpdatePost)
		// 评论的相关操作
		apiv1.POST("/post/:id/like", apiservice.LikePost)
		// 获取验证码
		apiv1.GET("/capacha", apiservice.GetCapacha)
		apiv1.POST("/capacha", apiservice.VerfiyCaptcha)
		// 发送邮件
		apiv1.POST("/email", apiservice.SendRegisterMail)
		// 上传图片
		apiv1.POST("/image/upload", apiservice.CkeditorUpload)
		apiv1.POST("/attach/upload", apiservice.UploadAttach)
		apiv1.POST("/attach/add", apiservice.UploadAddAttach)
		apiv1.POST("/attach/delete", apiservice.DeleteAttach)
	}

	// 管理员页面
	admin := r.Group("/admin")
	{
		// 登录展示页
		admin.GET("/login.html", adminservice.AdminLogin)
		// 管理员二次登录验证
		admin.POST("/login", adminservice.AdminLoginCheck)
		admin.GET("/setting/base.html", adminservice.AdminSettingBase)
		admin.POST("/setting/base", adminservice.AdminSettingBaseUpdate)
		admin.GET("/setting/smtp.html", adminservice.AdminSettingSmtp)
		admin.POST("/setting/smtp", adminservice.AdminSettingSmtpUpdate)
		admin.GET("/setting/extend.html", adminservice.AdminSettingExtend)
		admin.GET("/user/list.html", adminservice.AdminUserList)
		admin.GET("/user/group.html", adminservice.AdminGroupList)
		admin.GET("/user/create.html", adminservice.AdminUserCreate)
		admin.POST("/user/add", adminservice.AdminUserAdd)
	}
	admin.Use(jwt.JWT())
	{
		// 后台首页
		admin.GET("/index.html", adminservice.AdminIndex)
		// 模块列表
		admin.GET("/forum_list.html", adminservice.GetForumList)
		// 模块：新建
		admin.GET("/forum_new.html", adminservice.NewForum)
		admin.POST("/forum", adminservice.AddForum)
	}

	type Te struct {
		Name    string                 `json:"name"`
		TestArr []string               `json:"test_arr"`
		Love    map[string]interface{} `json:"love"`
	}
	testr := r.Group("/test")
	{

		// 检测session设置可用性
		testr.GET("/setsession", apiservice.TestSetSesssion)
		testr.GET("/getsession", apiservice.TestGetSesssion)
		testr.GET("/delsession", apiservice.TestDelSesssion)

		testr.GET("/strtime", func(c *gin.Context) {
			ress, _ := model.GetForumsList("id asc")
			//res := StrTime(1545793886)
			res := StrTime(ress[0].CreatedAt.Unix())
			c.JSON(200, gin.H{"res": res, "unic": ress[0].CreatedAt.Unix()})
		})

		testr.POST("/parm", func(c *gin.Context) {
			//name := c.PostForm("name")
			var tt Te
			err := c.ShouldBind(&tt)
			if err != nil {
				c.JSON(200, gin.H{"err": err.Error()})
				return
			}
			c.JSON(200, gin.H{"res": tt, "test": tt.Love})
		})

		testr.GET("/increment", func(c *gin.Context) {
			err := model.Increment("bbs_post", 1, "files_num")
			c.JSON(200, gin.H{"err": err})
		})

		testr.GET("cook/set", func(c *gin.Context) {

			err := package_redis.Set("uid1newthread1565420078659auzk201y2ve", "uid1newthread1565420078659auzk201y2ve", 100000)

			c.JSON(200, gin.H{"message": err})
		})

		testr.GET("cook/get", func(c *gin.Context) {
			value, err := package_redis.Get("no")
			c.JSON(200, gin.H{"value": value, "message": err.Error()})
		})
	}

	return r
}
