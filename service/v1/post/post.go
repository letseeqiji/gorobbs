package post

import (
	"gorobbs/model"
	"time"
)

// 回复帖子后的东西
/**
只有登录才可以回帖
	        -- forum       todayposts_cnt+1
	        -- thread      last_date  posts_cnt+1 first_post_id last_user_id last_post_id
	        -- post        插入一条数据:thread_id user_id isfirst=0 userip doctype=0 message
            -- mypost      插入一条数据:user_id thread_id post_id
            user        posts_cnt+1 credits_num+1[回帖帖加积分]
*/
func AfterAddNewPost(post *model.Post, threadID int) {
	threadInfo, _ := model.GetThreadById(threadID)
	forumID := threadInfo.ForumID

	oldFoumInfo, _ := model.GetForumByID(forumID)
	model.UpdateForumTodaypostsCnt(forumID, oldFoumInfo.TodaypostsCnt+1)

	updateThread := model.Thread{
		LastDate:   time.Now(),
		PostsCnt:   threadInfo.PostsCnt + 1,
		LastUserID: post.UserID,
		LastPostID: int(post.ID),
	}
	model.UpdateThread(threadID, updateThread)

	model.AddMyPost(post.UserID, threadID, int(post.ID))

	oldUserInfo, _ := model.GetUserByID(post.UserID)

	model.UpdateUserPostsCnt(post.UserID, oldUserInfo.PostsCnt+1)
	model.UpdateUserCreditsNum(post.UserID, oldUserInfo.CreditsNum+1)
}
