package util

import (
	"gorobbs/package/setting"
	"strconv"

	"gopkg.in/gomail.v2"
)

type Mail struct {
	MailTo  []string
	Subject string
	Body    string
}

var mailConn map[string]string

//var mail *Mail 谁什么错了

// 如何直接初始化类型中的值
func init() {
	user := setting.SmtpSetting.EmailUser
	pass := setting.SmtpSetting.EmailPass
	host := setting.SmtpSetting.EmailHost
	port := setting.SmtpSetting.EmailPort

	mailConn = map[string]string{
		"user": user,
		"pass": pass,
		"host": host,
		"port": port,
	}
}

func (mail *Mail) SendMail() error {
	//定义邮箱服务器连接信息，如果是阿里邮箱 pass填密码，qq邮箱填授权码
	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()
	m.SetHeader("From", "xinshulaile"+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("To", mail.MailTo...)                           //发送给多个用户
	m.SetHeader("Subject", mail.Subject)                        //设置邮件主题
	m.SetBody("text/html", mail.Body)                           //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	return err
}
