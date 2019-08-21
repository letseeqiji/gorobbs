package bootstrap

import (
	emailtool "gorobbs/tools/email"
)

func SetUp() {
	go emailtool.SendQueueEmail()
}
