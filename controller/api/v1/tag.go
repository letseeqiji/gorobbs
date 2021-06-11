package v1

import (
	"gorobbs/package/app"
	"gorobbs/package/rcode"
	string_package "gorobbs/package/string"
	"gorobbs/service/v1/tag"

	"github.com/gin-gonic/gin"
)

func AddTag(c *gin.Context) {
	cateid := string_package.A2i(c.DefaultPostForm("cate_id", ""))
	name := c.DefaultPostForm("name", "")
	comment := c.DefaultPostForm("comment", "")
	rank := string_package.A2i(c.DefaultPostForm("rank", "0"))
	enable := string_package.A2i(c.DefaultPostForm("enable", "1"))
	isdefault := string_package.A2i(c.DefaultPostForm("isdefault", "0"))
	style := c.DefaultPostForm("style", "secondary")
	code := rcode.SUCCESS

	if len(name) == 0 {
		code = rcode.INVALID_PARAMS
		app.JsonErrResponse(c, code)
		return
	}

	newTag, err := tag.AddTag(cateid, name, style, comment, rank, enable)
	if err != nil {
		code = rcode.ERROR_SQL_INSERT_FAIL
		app.JsonErrResponse(c, code)
		return
	}

	if isdefault == 1 && enable == 1 {
		tag.UpdateTagCateDefaultTagIDByID(cateid, newTag.ID)
	}

	app.JsonOkResponse(c, code, nil)
}

func UpdateTag(c *gin.Context) {
	id := string_package.A2i(c.DefaultPostForm("id", ""))
	cateid := string_package.A2i(c.DefaultPostForm("cate_id", ""))
	name := c.DefaultPostForm("name", "")
	comment := c.DefaultPostForm("comment", "")
	rank := string_package.A2i(c.DefaultPostForm("rank", "0"))
	enable := string_package.A2i(c.DefaultPostForm("enable", "1"))
	isdefault := string_package.A2i(c.DefaultPostForm("isdefault", "0"))
	style := c.DefaultPostForm("style", "secondary")
	code := rcode.SUCCESS

	if len(name) == 0 {
		code = rcode.INVALID_PARAMS
		app.JsonErrResponse(c, code)
		return
	}

	err := tag.UpdateTag(id, cateid, name, style, comment, rank, enable)
	if err != nil {
		code = rcode.ERROR_SQL_UPDATE_FAIL
		app.JsonErrResponse(c, code)
		return
	}

	// 启用并且设置为默认   如果启用并设置为非默认 怎么办呢？如果原来是默认的，设置默认为0  如果原来就不是默认的，就不用 操作了
	if isdefault == 1 && enable == 1 {
		tag.UpdateTagCateDefaultTagIDByID(cateid, id)
	}

	app.JsonOkResponse(c, code, nil)
}
