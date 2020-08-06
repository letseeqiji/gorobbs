package v1

import (
	"gorobbs/model"
	"gorobbs/package/app"
	"gorobbs/package/logging"
	"gorobbs/package/rcode"
	"gorobbs/package/session"
	tag_service "gorobbs/service/v1/tag"
	searchtool "gorobbs/tools/search"
	"gorobbs/tools/sensitivewall"
	"gorobbs/util"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	file_package "gorobbs/package/file"
	string_package "gorobbs/package/string"
	thread_service "gorobbs/service/v1/thread"
	user_service "gorobbs/service/v1/user"
)

// 添加帖子
func AddThread(c *gin.Context) {
	// 一致的信息：模块forum.id thread.threadname post.message 以及登录的用户信息

	// 获取forumid， docutype uid tname，postmessage
	// 验证 登录  所有字段不能为空
	// 添加记录:添加表thread 和 表 post

	forum_id, _ := strconv.Atoi(c.DefaultPostForm("forum_id", "1"))
	doctype, _ := strconv.Atoi(c.DefaultPostForm("doctype", "0"))
	subject := c.DefaultPostForm("subject", "")
	// 防止xss攻击
	subject = util.XssPolice(subject)
	subject, res := sensitivewall.Check(subject, "***")
	message := c.DefaultPostForm("message", "")
	// 防止xss攻击
	message = util.XssPolice(message)
	message, res = sensitivewall.Check(message, "***")

	// 如果包含非法字符
	if res {
		// TODO 帖子表设置一个标志位，需要人工审核帖子内容，而且显示默认不显示需要审核的帖子
		//app.JsonErrResponse(c, rcode.INVALID_CONTENT)
		//return
	}

	attachFileString := c.PostForm("attachfiles")
	tagThreadString := c.PostForm("tagthreads")
	attachfiles := []string{}
	tagThread := []string{}
	filesNum := 0
	code := rcode.SUCCESS

	if len(attachFileString) > 0 {
		attachfiles = strings.Split(attachFileString, ",")
		filesNum = len(attachfiles)
	}

	if len(tagThreadString) > 0 {
		tagThread = strings.Split(tagThreadString, ",")
	}

	uid, _ := strconv.Atoi(session.GetSession(c, "userid"))
	uip := c.ClientIP()

	thread := &model.Thread{
		ForumID:  forum_id,
		UserID:   uid,
		Userip:   uip,
		Subject:  subject,
		FilesNum: filesNum,
		LastDate: time.Now(),
	}

	newThread, err := model.AddThread(thread)
	if err != nil {
		logging.Info("thread入库错误", err.Error())
		code = rcode.ERROR_SQL_INSERT_FAIL
		app.JsonErrResponse(c, code)
		return
	}

	post := &model.Post{
		ThreadID:   int(newThread.ID),
		UserID:     uid,
		Isfirst:    1,
		Userip:     uip,
		Doctype:    doctype,
		Message:    message,
		MessageFmt: message,
	}
	newPost, err := model.AddPost(post)
	if err != nil {
		logging.Info("post入库错误", err.Error())
		code = rcode.ERROR
		code = rcode.ERROR_SQL_INSERT_FAIL
		app.JsonErrResponse(c, code)
		return
	}

	// 记录thread的firstpostid
	model.UpdateThread(int(newThread.ID), model.Thread{FirstPostID: int(newPost.ID), LastDate: time.Now()})

	// 已经添加完了帖子信息
	thread_service.AfterAddNewThread(newThread)

	// 添加搜索缓存
	searchtool.AddSearchIndex(uint64(newThread.ID), newThread.Subject)

	// 记录附件表
	if len(attachFileString) > 0 {
		for _, attachfile := range attachfiles {
			file := strings.Split(attachfile, "|")
			fname := file[0]
			forginname := file[1]
			ftype := file_package.GetType(fname)
			ofile, err := os.Open(fname)
			defer ofile.Close()
			if err != nil {
				continue
			}
			fsize, _ := file_package.GetSize(ofile)
			attach := &model.Attach{
				ThreadID:    int(newThread.ID),
				PostID:      int(newPost.ID),
				UserID:      uid,
				Filename:    fname,
				Orgfilename: forginname,
				Filetype:    ftype,
				Filesize:    fsize,
			}
			_, err = model.AddAttach(attach)
			if err != nil {
				logging.Info("attach入库错误", err.Error())
				code = rcode.ERROR_SQL_INSERT_FAIL
				app.JsonErrResponse(c, code)
				return
			}
		}
	}

	for _, tagid := range tagThread {
		tag_service.AddTagThread(string_package.A2i(tagid), int(newThread.ID), forum_id)
	}

	app.JsonOkResponse(c, code, nil)
}

type Tids struct {
	Tidarr []string `json:"tidarr"`
}

// 删除
func DeleteThreads(c *gin.Context) {
	ids := c.PostForm("tidarr")
	code := rcode.SUCCESS
	idsSlice := strings.Split(ids, ",")

	// 验证管理员才可以
	uid, _ := strconv.Atoi(session.GetSession(c, "userid"))
	isadmin := user_service.IsAdmin(uid)
	if isadmin == "0" {
		code = rcode.UNPASS
		app.JsonErrResponse(c, code)
		return
	}

	err := thread_service.DelThreads(idsSlice)
	if err != nil {
		code = rcode.ERROR
		app.JsonErrResponse(c, code)
		return
	}

	app.JsonOkResponse(c, code, ids)
}

// 移动
func MoveThreads(c *gin.Context) {
	ids := c.PostForm("tidarr")
	newfid, _ := strconv.Atoi(c.PostForm("newfid"))
	code := rcode.SUCCESS

	if err := model.UpdateThreadForum(ids, newfid); err != nil {
		code = rcode.ERROR_SQL_UPDATE_FAIL
		app.JsonErrResponse(c, code)
		return
	}

	app.JsonOkResponse(c, code, nil)
}

// 置顶
func TopThreads(c *gin.Context) {
	ids := c.PostForm("tidarr")
	top, _ := strconv.Atoi(c.PostForm("top"))

	threadIdArr := strings.Split(ids, ",")
	code := rcode.SUCCESS

	// 置顶
	if top != 0 {
		for _, threadId := range threadIdArr {
			threadIdInt, _ := strconv.Atoi(threadId)
			threadInfo, _ := model.GetThreadById(threadIdInt)

			_, err := model.UpdateThread(threadIdInt, model.Thread{Top: top})
			if err != nil {
				code = rcode.ERROR_SQL_UPDATE_FAIL
				app.JsonErrResponse(c, code)
				return
			}

			// 没有置顶过--新增
			// 新增topthread数据  修改thread-top = top
			if threadInfo.Top == 0 {
				_, err = model.AddThreadToTop(&model.ThreadTop{
					ThreadID: threadIdInt,
					ForumID:  threadInfo.ForumID,
					Top:      top,
				})
				if err != nil {
					code = rcode.ERROR_SQL_INSERT_FAIL
					app.JsonErrResponse(c, code)
					return
				}
			} else {
				// 已经置顶过--修改
				// 修改topthread-top = top  修改thread-top = top
				err = model.UpdateThreadTopTo(threadIdInt, top)
				if err != nil {
					code = rcode.ERROR_SQL_UPDATE_FAIL
					app.JsonErrResponse(c, code)
					return
				}
			}
		}
	} else {
		for _, threadId := range threadIdArr {
			threadIdInt, _ := strconv.Atoi(threadId)
			threadInfo, err := model.GetThreadById(threadIdInt)
			if err != nil {
				continue
			}
			// 取消置顶 top = 0
			// 没有置顶过
			// 不操作
			if threadInfo.Top == 0 {
				continue
			} else {
				// 置顶过
				// 删除topthread中数据  thread-top=0
				model.UpdateThreadTop(threadIdInt, top)
				model.DelThreadTopByTid(threadIdInt)
			}
		}
	}

	app.JsonOkResponse(c, code, nil)
}

// 关闭
func CloseThreads(c *gin.Context) {
	/*ids := &Tids{}
	err := c.Bind(&ids)*/

	ids := c.PostForm("tidarr")
	close := c.PostForm("close")

	c.JSON(200, gin.H{
		"code":    200,
		"message": "",
		"data":    ids,
		"close":   close,
	})
}

// 审核
func AuditedThread(c *gin.Context) {
	id := string_package.A2i(c.Param("id"))
	audited := string_package.A2i(c.DefaultPostForm("audited", "1"))
	code := rcode.SUCCESS

	err := thread_service.AuditedThread(id, audited)
	if err != nil {
		code = rcode.ERROR_SQL_UPDATE_FAIL
		app.JsonErrResponse(c, code)
		return
	}

	app.JsonOkResponse(c, code, nil)
}

// 更新主题内容
func UpdateThread(c *gin.Context) {

	forum_id, _ := strconv.Atoi(c.DefaultPostForm("forum_id", "1"))
	thread_id, _ := strconv.Atoi(c.Param("id"))
	post_id, _ := strconv.Atoi(c.DefaultPostForm("post_id", "1"))
	doctype, _ := strconv.Atoi(c.DefaultPostForm("doctype", "0"))
	subject := c.DefaultPostForm("subject", "")
	// 防止xss
	subject = util.XssPolice(subject)
	subject, res := sensitivewall.Check(subject, "***")
	message := c.DefaultPostForm("message", "")
	// 防止xss
	message = util.XssPolice(message)
	message, res = sensitivewall.Check(message, "***")
	// 包含非法的内容
	if res {
		// TODO 帖子表设置一个标志位，需要人工审核帖子内容，而且显示默认不显示需要审核的帖子
		//app.JsonErrResponse(c, rcode.INVALID_CONTENT)
		//return
	}

	uid, _ := strconv.Atoi(session.GetSession(c, "userid"))
	uip := c.ClientIP()
	code := rcode.SUCCESS

	// 找thread
	oldThread, err := model.GetThreadById(thread_id)
	if err != nil {
		code = rcode.ERROR_UNFIND_DATA
		app.JsonErrResponse(c, code)
		return
	}
	if oldThread.UserID != uid {
		code = rcode.UNPASS
		app.JsonErrResponse(c, code)
		return
	}
	// 找post
	oldPost, err := model.GetThreadFirstPostByTid(thread_id)
	if err != nil {
		code = rcode.ERROR_UNFIND_DATA
		app.JsonErrResponse(c, code)
		return
	}
	if int(oldPost.ID) != post_id {
		code = rcode.UNPASS
		app.JsonErrResponse(c, code)
		return
	}

	thread := model.Thread{
		ForumID: forum_id,
		Userip:  uip,
		Subject: subject,
	}
	model.UpdateThread(thread_id, thread)

	post := model.Post{
		Userip:  uip,
		Doctype: doctype,
		Message: message,
	}
	model.UpdatePost(post_id, post)

	app.JsonOkResponse(c, code, nil)
}

// 添加附件
// 直接添加到表中，因为以及各有了帖子  所以可以直接添加
func AddthreadAttach(c *gin.Context) {
	// 获取文件内容
	// 获取threadid postid uid
	// 修改thread表的files字段 + 1
	// 在attach表中添加一天新的记录
}

// 删除帖子的附件  知己额删除  提供好attach的id  就能删除
func DelthreadAttach(c *gin.Context) {
	// 删除数据内容  删除文件内容
	// 获取threadid
	// 修改thread表的files字段 - 1
	// 在attach表中直接删除记录
}
