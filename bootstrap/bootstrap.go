package bootstrap

import (
	emailtool "gorobbs/tools/email"
	searchtool "gorobbs/tools/search"
)

func SetUp() {
	go emailtool.SendQueueEmail()
	go searchtool.SearchInit()
}
