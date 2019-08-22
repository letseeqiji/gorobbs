package thread

import "gorobbs/model"

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
