package bootstrap

import (
	emailtool "gorobbs/tools/email"
	searchtool "gorobbs/tools/search"
	seneitivetool "gorobbs/tools/sensitivewall"
)

func SetUp() {
	go emailtool.SendQueueEmail()
	go searchtool.SearchInit()
	go seneitivetool.Trieinit()
}
