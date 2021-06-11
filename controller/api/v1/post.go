package v1

import (
	"gorobbs/model"
	"gorobbs/package/app"
	"gorobbs/package/logging"
	"gorobbs/package/rcode"
	"gorobbs/package/session"
	post_service "gorobbs/service/v1/post"
	"gorobbs/util"
	"os"
	"strconv"
	"strings"

	file_package "gorobbs/package/file"

	"github.com/gin-gonic/gin"
)

func AddPost(c *gin.Context) {
	//threadId, _ := strconv.Atoi(c.Param("id"))
	// 获取threadid docutype uid postmessage
	// 验证 登录  所有字段不能为空
	// 添加记录

	tid, _ := strconv.Atoi(c.DefaultPostForm("threadid", "1"))
	docutype, _ := strconv.Atoi(c.DefaultPostForm("doctuype", "0"))
	message := c.DefaultPostForm("message", "")
	// 防止xss攻击
	message = util.XssPolice(message)
	uid, _ := strconv.Atoi(session.GetSession(c, "userid"))
	uip := c.ClientIP()
	code := rcode.SUCCESS

	attachFileString := c.PostForm("attachfiles")
	attachfiles := []string{}
	filesNum := 0
	if len(attachFileString) > 0 {
		attachfiles = strings.Split(attachFileString, ",")
		filesNum = len(attachfiles)
	}

	post := &model.Post{
		ThreadID:   tid,
		UserID:     uid,
		Isfirst:    0,
		Userip:     uip,
		Doctype:    docutype,
		Message:    message,
		MessageFmt: message,
		FilesNum:   filesNum,
	}

	newPost, err := model.AddPost(post)
	if err != nil {
		logging.Info("回复帖子入库错误", err.Error())
		code = rcode.ERROR_SQL_INSERT_FAIL
		app.JsonErrResponse(c, code)
		return
	}

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
				ThreadID:    int(tid),
				PostID:      int(newPost.ID),
				UserID:      uid,
				Filename:    fname,
				Orgfilename: forginname,
				Filetype:    ftype,
				Filesize:    fsize,
			}
			model.AddAttach(attach)
		}
	}

	// 评论后数据统计
	post_service.AfterAddNewPost(newPost, tid)

	app.JsonOkResponse(c, code, nil)
}

// 更新评论内容
func UpdatePost(c *gin.Context) {

	post_id, _ := strconv.Atoi(c.DefaultPostForm("post_id", "1"))
	doctype, _ := strconv.Atoi(c.DefaultPostForm("doctype", "0"))
	message := c.DefaultPostForm("message", "")
	reason := c.DefaultPostForm("update_reason", "")
	uid, _ := strconv.Atoi(session.GetSession(c, "userid"))
	uip := c.ClientIP()
	code := rcode.SUCCESS

	// 找thread
	oldPost, err := model.GetPostById(post_id)
	if err != nil {
		code = rcode.ERROR_UNFIND_DATA
		app.JsonErrResponse(c, code)
		return
	}
	if oldPost.UserID != uid {
		code = rcode.UNPASS
		app.JsonErrResponse(c, code)
		return
	}

	post := model.Post{
		Userip:  uip,
		Doctype: doctype,
		Message: message,
	}
	model.UpdatePost(post_id, post)

	postUplog := &model.PostUpdateLog{
		PostID:     post_id,
		UserID:     uid,
		Reason:     reason,
		Message:    message,
		OldMessage: oldPost.Message,
	}
	model.AddPostUpdateLog(postUplog)

	app.JsonOkResponse(c, code, nil)
}

// 帖子点赞--取消点赞
func LikePost(c *gin.Context) {
	/**
	根据用户id和postid查询是否已经点过赞，
		如果没点赞
			post-likescnt+1 postlike插入一条新数据[uid,postid]
		如果已经点过赞
			post-likescnt-1 postlike删除数据[uid,postid]
		返回当前的操作是点赞还是取消点赞，post的likecnt
	*/
	postId, _ := strconv.Atoi(c.DefaultPostForm("postid", "0"))
	uid, _ := strconv.Atoi(session.GetSession(c, "userid"))

	postInfo, err := model.GetPostById(postId)
	if err != nil {
		c.JSON(404, gin.H{"code": 404, "message": err.Error()})
		return
	}

	// 获取是否收藏
	islike, _ := model.CheckPostlike(uid, postId)
	action := 0
	likesCnt := postInfo.LikesCnt

	// 如果没有点过赞 添加点赞
	if islike == 0 {
		model.AddPostlike(uid, postId)
		model.UpdatePostLikesNum(postId, postInfo.LikesCnt+1)
		action = 1
		likesCnt++
	} else {
		model.DelPostlike(uid, postId)
		if postInfo.LikesCnt > 0 {
			model.UpdateThreadFavouriteCnt(postId, postInfo.LikesCnt-1)
		}
		action = 0
		likesCnt--
	}

	c.JSON(200, gin.H{
		"code":      200,
		"message":   "ok",
		"action":    action,
		"likes_cnt": likesCnt,
	})
}
