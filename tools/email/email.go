package email

import (
	"encoding/json"
	"fmt"
	"gorobbs/package/queue"
	"gorobbs/package/setting"
	"gorobbs/service/v1/email"
	"strconv"
	"time"
)

// 读取redis中的邮件地址  发送邮件
func SendQueueEmail() {
	port := strconv.Itoa(setting.ServerSetting.HttpPort)
	webPre := setting.ServerSetting.Sitepre
	webUrl := setting.ServerSetting.Siteurl
	host := webPre + "://" + webUrl
	if setting.ServerSetting.RunMode == "debug" {
		host = host + ":" + port
	}

	emailTo, err := queue.Get(webUrl)
	var email2 string
	json.Unmarshal(emailTo, &email2)
	if err == nil {
		err = email.SendRegisterMail2(host, email2)
		fmt.Println("发送了一条:", email2)

		if err != nil {
			queue.Set(webUrl, email2)
			fmt.Println("发送err:", err.Error())
		}
	}

	time.AfterFunc(time.Second * 3, SendQueueEmail)
}
