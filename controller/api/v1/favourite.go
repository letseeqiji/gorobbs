package v1

import (
	"github.com/gin-gonic/gin"
	"gorobbs/model"
	"gorobbs/package/session"
	"strconv"
)

// 收藏--取消收藏
// 直接添加到表中，因为以及各有了帖子  所以可以直接添加
func Addthreadfavourite(c *gin.Context) {
	tid, _ := strconv.Atoi(c.DefaultPostForm("threadid", "1"))
	uid, _ := strconv.Atoi(session.GetSession(c, "userid"))

	threadInfo, err := model.GetThreadById(tid)
	if err != nil {
		c.JSON(404, gin.H{"code":404, "message":err.Error()})
		return
	}

	// 获取是否收藏
	isfav, _ := model.CheckFavourite(uid, tid)
	action := 0
	favNum := threadInfo.FavouriteCnt

	// 如果是1  添加收藏
	if isfav == 0 {
		model.AddMyFavourite(uid, tid)
		model.UpdateThreadFavouriteCnt(tid, threadInfo.FavouriteCnt+1)
		action = 1
		favNum ++
	} else {
		model.DelMyFavourite(uid, tid)
		if threadInfo.FavouriteCnt > 0 {
			model.UpdateThreadFavouriteCnt(tid, threadInfo.FavouriteCnt-1)
		}
		action = 0
		favNum --
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "ok",
		"action" : action,
		"fav_num" : favNum,
	})
}


