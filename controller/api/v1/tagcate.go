package v1

import (
	"gorobbs/package/app"
	"gorobbs/package/rcode"
	string_package "gorobbs/package/string"
	"gorobbs/service/v1/tag"

	"github.com/gin-gonic/gin"
)

func AddTagCate(c *gin.Context) {
	name := c.DefaultPostForm("name", "")
	comment := c.DefaultPostForm("comment", "")
	rank := string_package.A2i(c.DefaultPostForm("rank", "0"))
	enable := string_package.A2i(c.DefaultPostForm("enable", "1"))
	isforce := string_package.A2i(c.DefaultPostForm("isforce", "0"))
	style := c.DefaultPostForm("style", "secondary")
	code := rcode.SUCCESS

	if len(name) == 0 {
		code = rcode.INVALID_PARAMS
		app.JsonErrResponse(c, code)
		return
	}

	err := tag.AddTagCate(name, style, comment, rank, enable, isforce)
	if err != nil {
		code = rcode.ERROR_SQL_INSERT_FAIL
		app.JsonErrResponse(c, code)
		return
	}

	app.JsonOkResponse(c, code, nil)
}

func UpdateTagCate(c *gin.Context) {
	id := string_package.A2i(c.DefaultPostForm("id", ""))
	name := c.DefaultPostForm("name", "")
	comment := c.DefaultPostForm("comment", "")
	rank := string_package.A2i(c.DefaultPostForm("rank", "0"))
	enable := string_package.A2i(c.DefaultPostForm("enable", "1"))
	isforce := string_package.A2i(c.DefaultPostForm("isforce", "0"))
	defaultTagid := string_package.A2i(c.DefaultPostForm("default_tag_id", "0"))
	style := c.DefaultPostForm("style", "secondary")
	code := rcode.SUCCESS

	if len(name) == 0 {
		code = rcode.INVALID_PARAMS
		app.JsonErrResponse(c, code)
		return
	}

	err := tag.UpdateTagCate(id, name, style, comment, rank, enable, isforce, defaultTagid)
	if err != nil {
		code = rcode.ERROR_SQL_UPDATE_FAIL
		app.JsonErrResponse(c, code)
		return
	}

	app.JsonOkResponse(c, code, nil)
}
