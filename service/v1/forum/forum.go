package forum

import (
	"gorobbs/model"
	thread_service "gorobbs/service/v1/thread"
	"strconv"
)

const PAGE_SIZE int = 20

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

// 根据fid和page查询模块下的帖子列表  按照更新日期倒序排序
func DelForumByID(forumID int) (err error) {
	// 先获取该forum下的所有thread的id: select id from bbs_thread where forum_id = xx
	threads, err := model.GetThreadIDSByForumID(forumID)
	if err != nil {
		return err
	}

	var tids []string

	for _, v := range threads {
		tids = append(tids, strconv.Itoa(int(v.ID)))
	}

	// 删除帖子 及 相关数据  -- 这里应该分批处理 优化下
	// TODO 分批处理删除 防止一次性处理太多
	err = thread_service.DelThreads(tids)
	if err != nil {
		return
	}

	// 最后删除forum
	err = model.DelForumByID(forumID)

	return
}

// 更新forum
func UpdateForum(id, rank int, name, brief, announcement, moduids, seoTitle string) (err error) {
	whereMap := map[string]interface{}{"id": id}
	items := map[string]interface{}{"rank": rank, "name": name, "brief": brief, "announcement": announcement, "moduids": moduids, "seo_title": seoTitle}
	err = model.UpdateForum(whereMap, items)
	return
}
