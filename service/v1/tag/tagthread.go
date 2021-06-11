package tag

import "gorobbs/model"

/**
* @des 新增分类
* @param tagCate *TagCate 实例
* @return TagCate, error
 */
func AddTagThread(tagid, threadid, forumid int) (err error) {
	whereMap := model.TagThread{TagID: tagid, ThreadID: threadid, ForumID: forumid}
	model.AddTagThread(&whereMap)
	return
}
