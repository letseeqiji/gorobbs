package bootstrap

import (
	"gorobbs/controller/web"
	emailtool "gorobbs/tools/email"
)

func SetUp() {
	go emailtool.SendQueueEmail()
	go web.SearchInit()
}
