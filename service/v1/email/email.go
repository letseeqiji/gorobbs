package email

import (
	"fmt"
	"gorobbs/package/email"
	"gorobbs/package/setting"
	"gorobbs/util"
	"strconv"
	"time"
)

func SendRegisterMail(mailTo string) error {
	mail := util.Mail{}
	mail.MailTo = []string{
		mailTo,
	}

	webname := setting.ServerSetting.Sitename

	//邮件主题为"Hello"
	mail.Subject = "【" + webname + "】请验证您的邮件地址"
	//生成验证码
	randCode := util.GenRandCode(6)

	// todo 将验证码和邮箱存到全局空间  内存 redis session 等 都可以
	util.GlobalSet(mailTo, randCode)
	// 邮件正文
	mail.Body = "<h3>欢迎注册新书来了，验证码 " + randCode + "，5分钟内有效</h3>"
	err := mail.SendMail()

	return err
}

func SendRegisterMail2(host string, mailTo string) error {
	mail := util.Mail{}
	mail.MailTo = []string{
		setting.SmtpSetting.EmailUser,
		mailTo,
	}

	webname := setting.ServerSetting.Sitename
	now := strconv.Itoa(int(time.Now().Unix()))
	sign := util.EncodeMD5(mailTo + now)

	// redis记录
	email.Set(mailTo, now, 24*60*60)

	//邮件主题为"Hello"
	mail.Subject = "【" + webname + "】请验证您的邮件地址"

	href := "%s/register/checkMail?email=%s&time=%s&sign=%s"
	href = fmt.Sprintf(href, host, mailTo, now, sign)

	// 发送主题
	body := `尊敬的用户 <a data-auto-link="1" href="mailto:%s">%s</a>，您好：
	<p>
		您使用了邮箱 <a data-auto-link="1" href="mailto:%s">%s</a> 注册成为【%s】的会员。请点击以下链接，确认您在%s的注册：<br>
	<a href="%s" target="_blank">%s</a><br><br>

	如果以上链接不能点击，你可以复制网址URL，然后粘贴到浏览器地址栏打开，完成确认。<br><br>

		%s<br><br>

	（这是一封自动发送的邮件，请不要直接回复）<br><br>

		说明<br><br>

	－如果你没有注册过%s，可能是有人尝试使用你的邮件来注册，请忽略本邮件。<br>
	－没有激活的账号会为你保留24个小时, 请尽快激活。<br>
	－24个小时以后, 没有被激活的注册会自动失效，你需要重新填写并注册。<br>
	</p>`
	body = fmt.Sprintf(body, mailTo, mailTo, mailTo, mailTo, webname, webname, href, href, webname, webname)

	// 邮件正文
	mail.Body = body
	err := mail.SendMail()

	return err
}

func SendResetPasswordMail(host string, mailTo string) error {
	mail := util.Mail{}
	mail.MailTo = []string{
		setting.SmtpSetting.EmailUser,
		mailTo,
	}

	webname := setting.ServerSetting.Sitename
	now := strconv.Itoa(int(time.Now().Unix()))
	sign := util.EncodeMD5(mailTo + now)

	// redis记录
	email.Set(mailTo, now, 24*60*60)

	//邮件主题为"Hello"
	mail.Subject = "【" + webname + "】重设密码"

	href := "%s/password/reset.html?email=%s&time=%s&sign=%s"
	href = fmt.Sprintf(href, host, mailTo, now, sign)

	// 发送主题
	body := `尊敬的用户 <a data-auto-link="1" href="mailto:%s">%s</a>，您好：
	<p>
		您使用了邮箱 <a data-auto-link="1" href="mailto:%s">%s</a> 找回【%s】的会员。请点击以下链接，可以重设您的密码：<br>
	<a href="%s" target="_blank">%s</a><br><br>

	如果以上链接不能点击，你可以复制网址URL，然后粘贴到浏览器地址栏打开，完成确认。<br><br>

		%s<br><br>

	（这是一封自动发送的邮件，请不要直接回复）<br><br>
	</p>`
	body = fmt.Sprintf(body, mailTo, mailTo, mailTo, mailTo, webname, href, href, webname)

	// 邮件正文
	mail.Body = body
	err := mail.SendMail()

	return err
}
