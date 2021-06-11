package admin

import (
	"gorobbs/package/setting"
	"gorobbs/service/v1/user"

	"github.com/gin-gonic/gin"
)

func AdminSettingBase(c *gin.Context) {
	/*
		这个页面读取配置文件的对应内容放入到表单中
		提交后把对应的内容存储到配置文件的对应字段
		难度一星
	*/
	sessions := user.GetSessions(c)
	serverSetting := setting.ServerSetting
	c.HTML(200, "asetting_base.html", gin.H{"serverSetting": serverSetting, "sessions": sessions})
}

// 修改基础配置
// sitename=%E6%96%B0%E4%B9%A6%E6%9D%A5%E4%BA%862&sitebrief=%E7%83%82%E5%8%%E6&runlevel=5&user_create_on=1&user_create_email_on=0&user_resetpw_on=1&lang=zh-cn
func AdminSettingBaseUpdate(c *gin.Context) {
	sitename := c.PostForm("sitename")
	sitebrief := c.PostForm("sitebrief")
	runlevel := c.PostForm("runlevel")
	user_create_on := c.PostForm("user_create_on")
	user_create_email_on := c.PostForm("user_create_email_on")
	user_resetpw_on := c.PostForm("user_resetpw_on")
	lang := c.PostForm("lang")

	// 保存新设定的值
	setting.UpdateItemValue("server", "Sitename", sitename)
	setting.UpdateItemValue("server", "Sitebrief", sitebrief)
	setting.UpdateItemValue("server", "Runlevel", runlevel)
	setting.UpdateItemValue("server", "UserCreateEmailOn", user_create_email_on)
	setting.UpdateItemValue("server", "UserCreateOn", user_create_on)
	setting.UpdateItemValue("server", "UserResetpwOn", user_resetpw_on)
	setting.UpdateItemValue("server", "Lang", lang)

	c.JSON(200, gin.H{
		"code":    200,
		"message": "保存成功",
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

func AdminSettingSmtp(c *gin.Context) {
	smtpSetting := setting.SmtpSetting
	sessions := user.GetSessions(c)
	c.HTML(200, "asetting_smtp.html", gin.H{"smtpSetting": smtpSetting, "sessions": sessions})
}

//  email=wowiwo%40yeah.net&host=smtp.yeah.net&port=25&user=wowiwo%40yeah.net&pass=1qazxsw2
func AdminSettingSmtpUpdate(c *gin.Context) {
	host := c.PostForm("host")
	port := c.PostForm("port")
	user := c.PostForm("user")
	pass := c.PostForm("pass")

	// 保存新设定的值
	setting.UpdateItemValue("smtp", "EmailHost", host)
	setting.UpdateItemValue("smtp", "EmailPort", port)
	setting.UpdateItemValue("smtp", "EmailUser", user)
	setting.UpdateItemValue("smtp", "EmailPass", pass)

	c.JSON(200, gin.H{
		"code":    200,
		"message": "保存成功",
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

func AdminSettingExtend(c *gin.Context) {
	c.HTML(200, "asetting_extend.html", gin.H{})
}
