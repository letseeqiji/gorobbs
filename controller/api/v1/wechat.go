package v1

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/valyala/fasthttp"
	"gorobbs/package/setting"
	user_service "gorobbs/service/v1/user"
)

type wechatAtokenAndOpenid struct {
	access_token string
	openid       string
}

type wechatUser struct {
	openid     string
	nickname   string
	sex        string
	headimgurl string
	unionid    string
}

var (
	appID       = setting.WechatSetting.AppID
	appSecret   = setting.WechatSetting.AppSecret
	callBackURL = setting.WechatSetting.CallBackURL
)

/*func WechatLogin(c *gin.Context) {
	echostr := c.Query("echostr")

	c.String(200, echostr)
}*/

func WechatUserCheck(c *gin.Context) {
	//appID := setting.WechatSetting.AppID
	//appSecret := setting.WechatSetting.AppSecret
	//callBackURL := setting.WechatSetting.CallBackURL
	wcode := c.Query("code")
	url := "https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + appID + "&secret=" + appSecret + "&code=" + wcode + "&grant_type=authorization_code"

	status, resp, err := fasthttp.Get(nil, url)
	if err != nil {
		fmt.Println("请求失败:", err.Error())
		return
	}

	if status != fasthttp.StatusOK {
		fmt.Println("请求没有成功:", status)
		return
	}

	var wechatAAndid wechatAtokenAndOpenid
	err = json.Unmarshal(resp, &wechatAAndid)
	if err != nil {

	}
	accesstoken := wechatAAndid.access_token
	openid := wechatAAndid.openid

	url = "https://api.weixin.qq.com/sns/userinfo?access_token=" + accesstoken + "&openid=" + openid + "&lang=zh_CN"
	status, resp, err = fasthttp.Get(nil, url)
	if err != nil {
		fmt.Println("请求失败:", err.Error())
		return
	}

	if status != fasthttp.StatusOK {
		fmt.Println("请求没有成功:", status)
		return
	}

	var wechatUserInfo wechatUser
	err = json.Unmarshal(resp, &wechatUserInfo)
	if err != nil {

	}

	// 以上是获取微信用户信息的过程
	// 现在该处理本地业务了
	wechatUnionID := wechatUserInfo.unionid
	//根据wechatUnionID查找用户表
	user, err := user_service.GetUserBywWechatUnionID(wechatUnionID)

	// 数据库发生错误
	if err != nil {
		// do something
	}

	user_service.SetWechatSession(c, wechatUnionID)
	fmt.Println(wechatUnionID)

	// 如果没有找到信息
	if err == gorm.ErrRecordNotFound {
		//跳转到登录页面
		//if 已经注册过 {
		//	输入用户名，密码 登录并绑定
		//	成功后回到首页
		//} else {
		//	跳转到注册页 进行注册 完成后 进行绑定
		//	成功后回到首页
		//}
	} else {
		//进行登录操作，相当于输入了用户名密码的登录
		//返回上一个页面并刷新
		// 3，验证通过 生成token和session

		// 生成session  使nginx报502错误
		var sok chan int = make(chan int, 1)
		go user_service.LoginSession(c, user, sok)
		<-sok

		//跳转到首页
		c.Redirect(200, "/index.html")
	}

	/*if 找到了用户 {
		进行登录操作，相当于输入了用户名密码的登录
		返回上一个页面并刷新
	} else {
		跳转到登录页面
		if 已经注册过 {
			输入用户名，密码 登录并绑定
			成功后回到首页
		} else {
			跳转到注册页 进行注册 完成后 进行绑定
			成功后回到首页
		}
	}*/
}
