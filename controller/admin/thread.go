package admin

import (
	"github.com/gin-gonic/gin"
	"gorobbs/model"
	string_package "gorobbs/package/string"
)

func GetThreadList(c *gin.Context) {
	stype := c.DefaultQuery("type", "1")
	status := string_package.A2i(c.DefaultQuery("status", "0"))
	var threads []model.Thread

	// 下面做的就是根据提供的条件搜索需要的东西

	if stype == "1" {
		//select * from thread where post = 0
		threads, _ = model.GetThreads(map[string]interface{}{"isclosed": 0, "audited": status}, "created_at desc", 100, 1)
	}

	c.HTML(200, "athreadlist.html", gin.H{"threads": threads})
}
