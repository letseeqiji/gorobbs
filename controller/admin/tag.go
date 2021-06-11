package admin

import (
	"gorobbs/model"
	string_package "gorobbs/package/string"
	"gorobbs/service/v1/tag"
	"gorobbs/service/v1/user"

	"github.com/gin-gonic/gin"
)

func GetTagList(c *gin.Context) {
	sessions := user.GetSessions(c)
	taglist, _ := model.GetTagCateList()
	c.HTML(200, "tag_list.html", gin.H{"cate": taglist, "sessions": sessions})
}

func NewTagCate(c *gin.Context) {
	sessions := user.GetSessions(c)
	id := string_package.A2i(c.Query("id"))
	tagCate, _ := tag.GetTagCateByID(id)
	c.HTML(200, "tagcate_new.html", gin.H{"cate": tagCate, "sessions": sessions})
}

func EditTagCate(c *gin.Context) {
	sessions := user.GetSessions(c)
	id := string_package.A2i(c.Query("id"))
	tagCate, _ := tag.GetTagCateByID(id)
	c.HTML(200, "tagcate_edit.html", gin.H{"cate": tagCate, "sessions": sessions})
}

func NewTag(c *gin.Context) {
	id := string_package.A2i(c.Query("id"))
	tag, _ := tag.GetTagByID(id)
	taglist, _ := model.GetTagCateList()
	c.HTML(200, "tag_new.html", gin.H{"cate": taglist, "tag": tag})
}

func EditTag(c *gin.Context) {
	id := string_package.A2i(c.Query("id"))
	tag, _ := tag.GetTagByID(id)
	taglist, _ := model.GetTagCateList()
	c.HTML(200, "tag_edit.html", gin.H{"cate": taglist, "tag": tag})
}
