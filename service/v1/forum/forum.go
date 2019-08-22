package forum

import "gorobbs/model"

const PAGE_SIZE int  = 20

//新增分类
func AddForum(forumIcon string, forumName string, forumRank int) (forum *model.Forum, err error) {
	return model.AddForum("/"+forumIcon, forumName, forumRank)
}

// 根据fid和page查询模块下的帖子列表  按照更新日期倒序排序
func GetThreadListByForumID(forumID int, page int) (threads []model.Thread, err error) {
	threads, err = model.GetThreadListByForumID(forumID, page, PAGE_SIZE, "updated_at desc, last_date desc")
	return
}

// 根据fid和page查询模块下的帖子列表  按照更新日期倒序排序
func GetThreadTotleByForumID(forumID int) (num int) {
	//num = model.GetForumTotal("forum_id=?", forumID)
	num = model.GetThreadTotal(&model.Thread{ForumID: forumID})
	return
} 
