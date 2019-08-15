package v1

import (
	"gorobbs/model"
	"gorobbs/package/logging"
	"gorobbs/package/rcode"
	"gorobbs/package/session"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	thread_service "gorobbs/service/v1/thread"
	file_package "gorobbs/package/file"
)

// 添加帖子
func AddThread(c *gin.Context) {
	// 一致的信息：模块forum.id thread.threadname post.message 以及登录的用户信息

	// 获取forumid， docutype uid tname，postmessage
	// 验证 登录  所有字段不能为空
	// 添加记录:添加表thread 和 表 post

	/*doctype: "0"
	  forum_id: "2"
	  message: "<p>dddd</p>↵"
	  subject: "天津步履科技"
	*/
	forum_id, _ := strconv.Atoi(c.DefaultPostForm("forum_id", "1"))
	doctype, _ := strconv.Atoi(c.DefaultPostForm("doctype", "0"))
	subject := c.DefaultPostForm("subject", "")
	message := c.DefaultPostForm("message", "")
	attachFileString := c.PostForm("attachfiles")
	attachfiles := []string{}
	filesNum := 0

	if len(attachFileString) > 0 {
		attachfiles = strings.Split(attachFileString, ",")
		filesNum = len(attachfiles)
	}

	uid, _ := strconv.Atoi(session.GetSession(c, "userid"))
	uip := c.ClientIP()

	thread := &model.Thread{
		ForumID: forum_id,
		UserID:  uid,
		Userip:  uip,
		Subject: subject,
		FilesNum:filesNum,
		LastDate:time.Now(),
	}
	code := rcode.SUCCESS
	newThread, err := model.AddThread(thread)
	if err != nil {
		logging.Info("注册入库错误",err.Error())
		code = rcode.ERROR
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": code,
			"msg":  rcode.GetMessage(code),
		})
		return
	}

	post := &model.Post{
		ThreadID: int(newThread.ID),
		UserID:   uid,
		Isfirst:  1,
		Userip:   uip,
		Doctype:  doctype,
		Message:  message,
		MessageFmt:message,
	}
	newPost, err := model.AddPost(post)
	if err != nil {
		logging.Info("注册入库错误",err.Error())
		code = rcode.ERROR
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": code,
			"msg":  rcode.GetMessage(code),
		})
		return
	}

	// 记录thread的firstpostid
	model.UpdateThread(int(newThread.ID), model.Thread{FirstPostID:int(newPost.ID),LastDate:time.Now()})

	// 已经添加完了帖子信息
	thread_service.AfterAddNewThread(newThread)

	// 记录附件表
	/*
	ThreadID     int    `gorm:"default:0" json:"thread_id"`     //主题id
	PostID       int    `gorm:"default:0" json:"post_id"`       //帖子id
	UserID       int    `gorm:"default:0" json:"user_id"`       //用户id
	Filesize     int    `gorm:"default:0" json:"filesize"`      //文件尺寸，
	Width        int    `gorm:"default:0" json:"width"`         //width
	Height       int    `gorm:"default:0" json:"height"`        //
	Filename     string `gorm:"default:''" json:"filename"`     //文件名称，
	Orgfilename  string `gorm:"default:''" json:"orgfilename"`  //上传的原文
	Filetype     string `gorm:"default:''" json:"filetype"`     //image
	Comment      string `gorm:"default:''" json:"comment"`      //文件注释
	DownloadsNum int    `gorm:"default:0" json:"downloads_num"` //下载次数
	CreditsNum   int    `gorm:"default:0" json:"credits_num"`   //需要的积分
	GoldsNum     int    `gorm:"default:0" json:"golds_num"`     //需要的金币
	RmbsNum      int    `gorm:"default:0" json:"rmbs_num"`      //需要的人民
	Isimage      int    `gorm:"default:0" json:"isimage"`       //是否为图片
	*/
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
				logging.Info("注册入库错误",err.Error())
				code = rcode.ERROR
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": code,
					"msg":  rcode.GetMessage(code),
				})
				return
			}
		}
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "ok",
	})
}

type Tids struct {
	Tidarr []string `json:"tidarr"`
}

func DeleteThreads(c *gin.Context) {
	/*ids := &Tids{}
	err := c.Bind(&ids)*/

	ids := c.PostForm("tidarr")

	if err := model.DelThread(ids); err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "",
		"data":    ids,
	})
}

func MoveThreads(c *gin.Context) {
	/*ids := &Tids{}
	err := c.Bind(&ids)*/

	ids := c.PostForm("tidarr")
	newfid, _ := strconv.Atoi(c.PostForm("newfid"))

	if err := model.UpdateThreadForum(ids, newfid); err != nil {
		c.JSON(200, gin.H{
			"code":    500,
			"message": "数据修改失败",
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "修改成功",
		/*"data":    ids,
		"newfid":  newfid,*/
	})
}

func TopThreads(c *gin.Context) {
	ids := c.PostForm("tidarr")
	top, _ := strconv.Atoi(c.PostForm("top"))

	threadIdArr := strings.Split(ids, ",")

	// 置顶
	if top != 0 {
		for _, threadId := range threadIdArr {
			threadIdInt, _ := strconv.Atoi(threadId)
			threadInfo, _ := model.GetThreadById(threadIdInt)

			model.UpdateThread(threadIdInt, model.Thread{Top:top})

			// 没有置顶过--新增
			// 新增topthread数据  修改thread-top = top
			if threadInfo.Top == 0 {
				model.AddThreadToTop(&model.ThreadTop{
					ThreadID:threadIdInt,
					ForumID: threadInfo.ForumID,
					Top:top,
				})
			} else {
				// 已经置顶过--修改
				// 修改topthread-top = top  修改thread-top = top
				model.UpdateThreadTopTo(threadIdInt, top)
			}
		}
	} else {
		for _, threadId := range threadIdArr {
			threadIdInt, _ := strconv.Atoi(threadId)
			threadInfo, _ := model.GetThreadById(threadIdInt)
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

	c.JSON(200, gin.H{
		"code":    200,
		"message": "ok",
	})
}

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

// 更新主题内容
func UpdateThread(c *gin.Context) {

	forum_id, _ := strconv.Atoi(c.DefaultPostForm("forum_id", "1"))
	thread_id, _ := strconv.Atoi(c.Param("id"))
	post_id, _ := strconv.Atoi(c.DefaultPostForm("post_id", "1"))
	doctype, _ := strconv.Atoi(c.DefaultPostForm("doctype", "0"))
	subject := c.DefaultPostForm("subject", "")
	message := c.DefaultPostForm("message", "")
	uid, _ := strconv.Atoi(session.GetSession(c, "userid"))
	uip := c.ClientIP()

	// 找thread
	oldThread, err := model.GetThreadById(thread_id)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}
	if oldThread.UserID != uid {
		c.JSON(200, gin.H{
			"code":    403,
			"message": "这必是你的帖子，无权操作",
		})
		return
	}
	// 找post
	oldPost, err := model.GetThreadFirstPostByTid(thread_id)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}
	if int(oldPost.ID) != post_id {
		c.JSON(200, gin.H{
			"code":    403,
			"message": "传值与数据不符合",
		})
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

	c.JSON(200, gin.H{
		"code":    200,
		"message": "更新成功",
	})
}

// 添加附件
// 直接添加到表中，因为以及各有了帖子  所以可以直接添加
func AddthreadAttach(c *gin.Context)  {
	// 获取文件内容
	// 获取threadid postid uid
	// 修改thread表的files字段 + 1
	// 在attach表中添加一天新的记录
}

// 删除帖子的附件  知己额删除  提供好attach的id  就能删除
func DelthreadAttach(c *gin.Context)  {
	// 删除数据内容  删除文件内容
	// 获取threadid
	// 修改thread表的files字段 - 1
	// 在attach表中直接删除记录
}

