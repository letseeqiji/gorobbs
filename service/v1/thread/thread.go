package thread

import (
	"gorobbs/model"
)

// 获取指定用户的最新的10条thread [id subject]
func GetUserThreads(uid int) (threads []model.Thread, err error) {
	whereMap := &model.Thread{UserID: uid}
	order := "created_at desc"
	limit := 10
	threads, err = model.GetThreads(whereMap, order, limit, 1)
	return
}

// 发信帖子后要增加一些统计信息
/*
发帖后影响的逻辑
	        -- forum       threads_cnt+1 todaythreads_cnt+1
	        -- thread      插入一条数据:forum_id  user_id userip subject
	        -- post        插入一条数据:thread_id user_id isfirst=1 userip doctype=0 message
	        -- mythread    插入一条数据:user_id thread_id
	        -- user        threads_cnt+1 credits_num+3[发帖加积分]
*/
func AfterAddNewThread(thread *model.Thread) {
	forumID := thread.ForumID
	threadID := thread.ID
	userID := thread.UserID

	oldFoumInfo, _ := model.GetForumByID(forumID)
	// forum 表的 ThreadsCnt  TodaythreadsCnt 字段增加1
	model.UpdateForumThreadsCnt(forumID, oldFoumInfo.ThreadsCnt+1)
	model.UpdateForumTodaythreadsCnt(forumID, oldFoumInfo.TodaythreadsCnt+1)

	// my_thread 自家一条信息
	model.AddMyThread(userID, int(threadID))

	// 用户表user ThreadsCnt+1 CreditsNum+3
	oldUserInfo, _ := model.GetUserByID(thread.UserID)
	model.UpdateUserThreadsCnt(userID, oldUserInfo.ThreadsCnt+1)
	model.UpdateUserCreditsNum(userID, oldUserInfo.CreditsNum+3)
}

func DelThreads(tids []string) (err error) {

	// 删除所有评论post
	err = model.DelPostsOfThread(tids)
	if err != nil {
		return
	}

	// 删除所有 置顶 threadtop
	err = model.DelthreadTopsOfThread(tids)
	if err != nil {
		return
	}

	// 删除所有 mythread
	err = model.DelMyThreadsOfThread(tids)
	if err != nil {
		return
	}

	// 删除所有 mypost
	err = model.DelMyPostsOfThread(tids)
	if err != nil {
		return
	}

	// 删除所有 myfavourite
	err = model.DelMyFavouritesOfThread(tids)
	if err != nil {
		return
	}

	// 删除所有 附件
	err = model.DelAttachsOfThread(tids)
	if err != nil {
		return
	}

	// 删除所有的thread
	err = model.DelThread(tids)
	if err != nil {
		return
	}

	return
}

// 给定threadid数组 查询
func GetThreadsByIDs(tidArr []string) (threads []*model.Thread, err error) {
	threads, err = model.GetThreadsByIDs(tidArr)
	return
}

// 审核
func AuditedThread(threadID, audited int) (err error) {
	items := map[string]interface{}{"audited": audited}
	_, err = model.UpdateThreadPro(threadID, items)
	return
}
