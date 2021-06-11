package v1

import (
	"gorobbs/model"
	"gorobbs/package/app"
	"gorobbs/package/rcode"
	"gorobbs/package/session"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 收藏--取消收藏
// 直接添加到表中，因为以及各有了帖子  所以可以直接添加
func Addthreadfavourite(c *gin.Context) {
	tid, _ := strconv.Atoi(c.DefaultPostForm("threadid", "1"))
	uid, _ := strconv.Atoi(session.GetSession(c, "userid"))
	code := rcode.SUCCESS

	threadInfo, err := model.GetThreadById(tid)
	if err != nil {
		code = rcode.ERROR_UNFIND_DATA
		app.JsonErrResponse(c, code)
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
		favNum++
	} else {
		model.DelMyFavourite(uid, tid)
		if threadInfo.FavouriteCnt > 0 {
			model.UpdateThreadFavouriteCnt(tid, threadInfo.FavouriteCnt-1)
		}
		action = 0
		favNum--
	}

	app.JsonOkResponse(c, code, map[string]interface{}{"action": action, "fav_num": favNum})
}
