package admin

import (
	"gorobbs/model"
	"gorobbs/package/app"
	"gorobbs/package/file"
	"gorobbs/package/rcode"
	"gorobbs/package/upload"
	forum_service "gorobbs/service/v1/forum"
	tag_service "gorobbs/service/v1/tag"
	"gorobbs/service/v1/user"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetForumList(c *gin.Context) {
	// 首先列出已经具有的模块
	forums, _ := model.GetForumsList("id asc")

	fnum := len(forums) + 1
	sessions := user.GetSessions(c)

	c.HTML(200, "aforumlist.html", gin.H{
		"forums":   forums,
		"fnum":     fnum,
		"sessions": sessions,
	})
}

func NewForum(c *gin.Context) {
	sessions := user.GetSessions(c)
	c.HTML(200, "aforum_new.html", gin.H{"sessions": sessions})
}

func AddForum(c *gin.Context) {
	icon, err := c.FormFile("forum_icon")
	fullFileName := "static/img/forum.png"
	code := rcode.SUCCESS

	if err == nil {
		fileName := icon.Filename
		// 限制图片的格式 和 大小
		if !upload.CheckImageExt(fileName) {
			code = rcode.ERROR_IMAGE_BAD_EXT
			app.JsonErrResponse(c, code)
			return
		}

		if !upload.CheckImageSize2(icon) {
			code = rcode.ERROR_IMAGE_TOO_LARGE
			app.JsonErrResponse(c, code)
			return
		}

		filePath := "upload/forum"
		// 判断路径是否存在 不存在则创建
		filePath, err = file.CreatePathInToday(filePath)
		if err != nil {
			code = rcode.ERROR_FILE_CREATE_FAIL
			app.JsonErrResponse(c, code)
			return
		}

		fullFileName = filePath + "/" + fileName
		err = c.SaveUploadedFile(icon, fullFileName)
		if err != nil {
			code = rcode.ERROR_FILE_SAVE_FAIL
			app.JsonErrResponse(c, code)
			return
		}
	}

	name := c.PostForm("forum_name")
	if len([]rune(name)) < 2 {
		c.JSON(200, gin.H{
			"code":    400,
			"message": "名字至少包含2个字符",
		})
		return
	}

	rankk := c.PostForm("forum_rank")
	rank, _ := strconv.Atoi(rankk)
	if rank < 0 {
		c.JSON(200, gin.H{
			"code":    400,
			"message": "请设置有效的排序",
		})
		return
	}

	// 入库
	_, err = forum_service.AddForum(fullFileName, name, rank)
	if err != nil {
		code = rcode.ERROR_SQL_INSERT_FAIL
		app.JsonErrResponse(c, code)
		return
	}

	app.JsonOkResponse(c, code, nil)
}

func DelForum(c *gin.Context) {
	fid, _ := strconv.Atoi(c.PostForm("forumid"))
	code := rcode.SUCCESS
	err := forum_service.DelForumByID(fid)
	if err != nil {
		code = rcode.ERROR_SQL_DELETE_FAIL
		app.JsonErrResponse(c, code)
		return
	}

	app.JsonOkResponse(c, code, nil)
}

type Forum struct {
	ForumId   string `json:"forum_id"`
	ForumIcon string `json:"forum_icon"`
	ForumName string `json:"forum_name"`
	ForumRank string `json:"forum_rank"`
}

type ForumsLst struct {
	Forums []Forum `json:"forums"`
	//Forums []string `json:forums`
}

func AddForum2(c *gin.Context) {
	/*{
		forums:[
			{id:1, name:23, rank:2, icon:23},
		]
	}*/
	var Forms ForumsLst
	err := c.ShouldBindJSON(&Forms)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	headers := make(map[string]interface{})
	headers["Content-Type"] = c.GetHeader("Content-Type")
	headers["receieve-data"] = Forms
	c.JSON(200, gin.H{
		"code":    200,
		"message": headers,
	})
}

func EditForum(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	usetags, _ := tag_service.GetTagForumsByForumID(id)
	sessions := user.GetSessions(c)
	taglist, _ := model.GetTagCateList()
	foruminfo, _ := model.GetForumByID(id)
	c.HTML(
		200,
		"aforum_edit.html",
		gin.H{
			"sessions": sessions,
			"forum":    foruminfo,
			"tagcate":  taglist,
			"usetags":  usetags,
		})
}

func UpdateForum(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	name := c.DefaultPostForm("name", "")
	rank, _ := strconv.Atoi(c.DefaultPostForm("rank", "0"))
	brief := c.DefaultPostForm("brief", "")
	announcement := c.DefaultPostForm("announcement", "")
	moduids := c.DefaultPostForm("moduids", "")
	seo_title := c.DefaultPostForm("seo_title", "")
	tidarr := c.DefaultPostForm("tidarr", "")
	code := rcode.SUCCESS

	err := forum_service.UpdateForum(id, rank, name, brief, announcement, moduids, seo_title)
	if err != nil {
		code = rcode.ERROR_SQL_UPDATE_FAIL
		app.JsonErrResponse(c, code)
		return
	}

	// TODO 需要优化--这样做效率太低，不过好处是可以保证数据紧邻
	tag_service.DeleteTagForum(id)
	if len(tidarr) != 0 {
		tids := strings.Split(tidarr, ",")
		for _, v := range tids {
			k, _ := strconv.Atoi(v)
			tag_service.AddTagForum(k, id)
		}
	}

	app.JsonOkResponse(c, code, nil)
}
