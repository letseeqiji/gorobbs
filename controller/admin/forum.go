package admin

import (
	"gorobbs/model"
	"gorobbs/package/file"
	"gorobbs/package/upload"
	forum_service "gorobbs/service/v1/forum"
	"gorobbs/service/v1/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetForumList(c *gin.Context) {
	// 首先列出已经具有的模块
	forums, _ := model.GetForumsList("id asc")

	fnum := len(forums) + 1
	sessions := user.GetSessions(c)

	c.HTML(200, "aforumlist.html", gin.H{
		"forums": forums,
		"fnum":   fnum,
		"sessions":sessions,
	})
}

func NewForum(c *gin.Context) {
	sessions := user.GetSessions(c)
	c.HTML(200, "aforum_new.html", gin.H{"sessions":sessions})
}

func AddForum(c *gin.Context) {
	icon, err := c.FormFile("forum_icon")
	if err != nil {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "上传错误",
		})
		return
	}



	fileName := icon.Filename
	// 限制图片的格式 和 大小
	if ! upload.CheckImageExt(fileName)  {
		c.JSON(200, gin.H{
			"code":    403,
			"message": "图片格式不正确",
		})
		return
	}

	if !upload.CheckImageSize2(icon) {
		c.JSON(200, gin.H{
			"code":    403,
			"message": "图片大小超标了",
		})
		return
	}

	filePath := "upload/forum"
	// 判断路径是否存在 不存在则创建
	filePath, err = file.CreatePathInToday(filePath)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	fullFileName := filePath + "/" + fileName
	err = c.SaveUploadedFile(icon, fullFileName)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
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
	newForum, err := forum_service.AddForum(fullFileName, name, rank)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "ok",
		"file":    newForum,
	})
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

	//c.Request.ParseForm()
	//	//Forms := c.Request.PostForm

	headers := make(map[string]interface{})
	headers["Content-Type"] = c.GetHeader("Content-Type")
	headers["receieve-data"] = Forms
	c.JSON(200, gin.H{
		"code":    200,
		"message": headers,
	})
}
