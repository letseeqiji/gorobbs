package email

import (
	"gorobbs/util"
)

func SendRegisterMail(mailTo string) error {
	mail := util.Mail{}
	mail.MailTo = []string{
		mailTo,
	}
	//邮件主题为"Hello"
	mail.Subject = "注册验证码"
	//生成验证码
	randCode := util.GenRandCode(6)

	// todo 将验证码和邮箱存到全局空间  内存 redis session 等 都可以
	util.GlobalSet(mailTo, randCode)
	// 邮件正文
	mail.Body = "<h3>欢迎注册新书来了，验证码 " + randCode + "，5分钟内有效</h3>"
	err := mail.SendMail()

	return err
}
