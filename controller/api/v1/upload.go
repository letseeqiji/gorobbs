package v1

import (
	"gorobbs/model"
	"gorobbs/package/app"
	file_package "gorobbs/package/file"
	"gorobbs/package/rcode"
	"gorobbs/package/session"
	"gorobbs/package/upload"
	"strconv"

	"github.com/gin-gonic/gin"
)

var pixgif = "data:image/gif;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVQImWNgYGBgAAAABQABh6FO1AAAAABJRU5ErkJggg=="

func CkeditorUpload(c *gin.Context) {
	file, _ := c.FormFile("upload")
	userid := session.GetSession(c, "userid")
	fileName := file.Filename

	if !upload.CheckImageSize2(file) {
		c.JSON(200, gin.H{
			"fileName": fileName,
			"uploaded": 1,
			"url":      pixgif,
		})
		return
	}

	newFilename := file_package.MakeFileName(userid, fileName)
	filepath := "upload/thread/" + userid
	filepath, err := file_package.CreatePathInToday(filepath)

	if err != nil {
		c.JSON(200, gin.H{
			"fileName": fileName,
			"uploaded": 1,
			"url":      pixgif,
		})
		return
	}

	fullName := filepath + "/" + newFilename
	// 上传文件到指定的路径
	c.SaveUploadedFile(file, fullName)

	c.JSON(200, gin.H{
		"fileName": fileName,
		"uploaded": 1,
		"url":      "/" + fullName,
	})
}

func UploadFile(c *gin.Context) {
	action := c.Query("action")
	uid := c.Query("uid")
	code := rcode.SUCCESS

	file, _ := c.FormFile("upload")
	fileName := file.Filename
	newFilename := file_package.MakeFileName(uid, fileName)
	if !upload.CheckImageSize2(file) {
		code = rcode.ERROR_IMAGE_TOO_LARGE
		app.JsonErrResponse(c, code)
		return
	}

	filepath := "upload/" + action + "/" + uid + "/"
	err := file_package.CreatePath(filepath)
	if err != nil {
		code = rcode.ERROR_FILE_CREATE_FAIL
		app.JsonErrResponse(c, code)
		return
	}

	fullName := filepath + newFilename
	// 上传文件到指定的路径
	err = c.SaveUploadedFile(file, fullName)
	if err != nil {
		code = rcode.ERROR_FILE_SAVE_FAIL
		app.JsonErrResponse(c, code)
		return
	}

	c.JSON(200, gin.H{"filename": "测试图", "filetype": 1, "url": "/" + fullName, "attatchid": 99})
}

/**
上传附件
	thread表的images和files
	attach表中每个文件一条记录
*/
func UploadAttach(c *gin.Context) {
	userid := session.GetSession(c, "userid")
	code := rcode.SUCCESS
	file, _ := c.FormFile("upload")
	fileName := file.Filename
	fileType := file_package.GetType(fileName)
	newFilename := file_package.MakeFileName(userid, fileName)
	if !upload.CheckImageSize2(file) {
		code = rcode.ERROR_IMAGE_TOO_LARGE
		app.JsonErrResponse(c, code)
		return
	}

	filepath := "upload/attach/" + userid
	filepath, err := file_package.CreatePathInToday(filepath)
	if err != nil {
		code = rcode.ERROR_FILE_CREATE_FAIL
		app.JsonErrResponse(c, code)
		return
	}

	fullName := filepath + "/" + newFilename
	// 上传文件到指定的路径
	err = c.SaveUploadedFile(file, fullName)
	if err != nil {
		code = rcode.ERROR_FILE_SAVE_FAIL
		app.JsonErrResponse(c, code)
		return
	}

	data := map[string]interface{}{"orgfilename": fileName, "filetype": fileType, "url": fullName}
	app.JsonOkResponse(c, code, data)
}

// 添加额外的附件
func UploadAddAttach(c *gin.Context) {
	userid := session.GetSession(c, "userid")
	threadId, _ := strconv.Atoi(c.DefaultPostForm("thread_id", "0"))
	posthreadId := threadId
	postId, _ := strconv.Atoi(c.PostForm("post_id"))
	code := rcode.SUCCESS

	file, _ := c.FormFile("upload")
	fileName := file.Filename
	fileType := file_package.GetType(fileName)
	fileSize := file.Size
	newFilename := file_package.MakeFileName(userid, fileName)
	if !upload.CheckImageSize2(file) {
		code = rcode.ERROR_IMAGE_TOO_LARGE
		app.JsonErrResponse(c, code)
		return
	}

	filepath := "upload/attach/" + userid
	filepath, err := file_package.CreatePathInToday(filepath)
	if err != nil {
		code = rcode.ERROR_FILE_CREATE_FAIL
		app.JsonErrResponse(c, code)
		return
	}

	fullName := filepath + "/" + newFilename
	// 上传文件到指定的路径
	err = c.SaveUploadedFile(file, fullName)
	if err != nil {
		code = rcode.ERROR_FILE_SAVE_FAIL
		app.JsonErrResponse(c, code)
		return
	}

	postInfo, _ := model.GetPostById(postId)
	if threadId == 0 {
		threadId = postInfo.ThreadID
	}

	useridInt, _ := strconv.Atoi(userid)
	// 不同发生在此处  由于是添加的  所以直接入库 修改thread的filesnum字段加1  返回attach名字id等信息
	model.AddAttach(&model.Attach{
		ThreadID:    threadId,
		PostID:      postId,
		UserID:      useridInt,
		Filesize:    int(fileSize),
		Filename:    fullName,
		Orgfilename: fileName,
		Filetype:    fileType,
	})

	// 评论添加时 不穿值
	if posthreadId != 0 {
		threadInfo, _ := model.GetThreadById(threadId)
		model.UpdateThreadFilesNum(threadId, threadInfo.FilesNum+1)
	}

	// 更新post表的filenum
	model.UpdatePostFilesNum(postId, postInfo.FilesNum+1)

	data := map[string]interface{}{"orgfilename": fileName, "filetype": fileType, "url": fullName}
	app.JsonOkResponse(c, code, data)
}

// 删除的附件
func DeleteAttach(c *gin.Context) {
	userid := session.GetSession(c, "userid")
	_ = userid
	attachId, _ := strconv.Atoi(c.PostForm("attach_id"))
	threadId, _ := strconv.Atoi(c.DefaultPostForm("thread_id", "0"))
	postId, _ := strconv.Atoi(c.DefaultPostForm("post_id", "0"))
	code := rcode.SUCCESS

	if threadId != 0 {
		threadInfo, _ := model.GetThreadById(threadId)
		if threadInfo.FilesNum != 0 {
			model.UpdateThreadFilesNum(threadId, threadInfo.FilesNum-1)
		}
		postId = threadInfo.FirstPostID
	}

	if postId != 0 {
		postInfo, _ := model.GetPostById(postId)
		if postInfo.FilesNum != 0 {
			model.UpdatePostFilesNum(postId, postInfo.FilesNum-1)
		}
	}

	model.DelAttach(attachId)

	app.JsonOkResponse(c, code, nil)
}
