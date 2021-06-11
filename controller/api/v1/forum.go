package v1

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/fatih/set.v0"
	"gorobbs/model"
	"gorobbs/package/app"
	"gorobbs/package/rcode"
	string_package "gorobbs/package/string"
	tag_service "gorobbs/service/v1/tag"
	"gorobbs/service/v1/thread"
	"strings"
)

func GetTagCateByForumID(c *gin.Context) {
	id := string_package.A2i(c.Param("id"))

	res, _ := tag_service.GetTagForumsByForumID(id)

	code := rcode.SUCCESS

	app.JsonOkResponse(c, code, res)
}

// 获取标签指示的帖子
func GetTagThreadsByForumIDWithTags(c *gin.Context) {
	id := string_package.A2i(c.Param("id"))
	tags := c.Query("tags")
	code := rcode.SUCCESS

	tagarr := strings.Split(tags, ",")
	ftag := string_package.A2i(tagarr[0])
	ltags := tagarr[1:]
	threadids, _ := model.GetTagThreadIDByForumIDWithTagID(id, ftag)
	var tidArr = set.New(set.ThreadSafe)
	for _, tagThread := range threadids {
		tidArr.Add(tagThread.ThreadID)
	}

	for _, tagid := range ltags {
		if !tidArr.IsEmpty() {
			threadids, _ := model.GetTagThreadIDByForumIDWithTagID(id, string_package.A2i(tagid))
			tidArr2 := set.New(set.ThreadSafe)
			for _, tagThread := range threadids {
				tidArr2.Add(tagThread.ThreadID)
			}

			tidArr = set.Intersection(tidArr, tidArr2)
			tidArr2.Clear()
		} else {
			break
		}
	}

	var threadidArr []string
	for {
		if !tidArr.IsEmpty() {
			threadidArr = append(threadidArr, string_package.I2A(tidArr.Pop().(int)))
		} else {
			break
		}
	}

	threads, _ := thread.GetThreadsByIDs(threadidArr)
	app.JsonOkResponse(c, code, threads)
}
