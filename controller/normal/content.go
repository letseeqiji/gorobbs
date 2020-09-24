package normal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorobbs/package/app"
	"gorobbs/package/rcode"
	"gorobbs/tools/sensitivewall"
	"strings"
)

type Content struct {
	Cont string `json:"content"`
}

// 检测内容是否包含非法信息  包含返回406  不包含返回201
func ContentCheck(c *gin.Context) {
	//var content Content

	content := c.DefaultPostForm("content", "")

	fmt.Println("收到的验证内容是:", content)

	content = strings.Replace(content, " ", "", -1)
	// 去除换行符
	content = strings.Replace(content, "\n", "", -1)
	content = strings.Replace(content, "&nbsp;", "", -1)
	content = strings.Replace(content, "<p>", "", -1)
	content = strings.Replace(content, "</p>", "", -1)
	content = strings.Replace(content, "#", "", -1)
	content = strings.Replace(content, "*", "", -1)
	content = strings.Replace(content, "&", "", -1)
	content = strings.Replace(content, "%", "", -1)
	content = strings.Replace(content, "$", "", -1)
	content = strings.Replace(content, "@", "", -1)
	content = strings.Replace(content, "(", "", -1)
	content = strings.Replace(content, ")", "", -1)
	fmt.Println("去掉空格收到的验证内容是:", content)

	if len(content) == 0 {
		app.JsonOkResponse(c, 201, nil)
	}

	illContent, res := sensitivewall.Check(content, "***")

	if res {
		fmt.Println("替换后的文本：", illContent)
		app.JsonErrResponse(c, rcode.INVALID_CONTENT)
		return
	}

	app.JsonOkResponse(c, 201, nil)
}
