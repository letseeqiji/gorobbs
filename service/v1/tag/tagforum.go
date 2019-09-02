package tag

import "gorobbs/model"

/**
* @des 新增分类
* @param tagCate *TagCate 实例
* @return TagCate, error
 */
func AddTagForum(tagcateid, forumid int) (err error) {
	whereMap := model.TagForum{TagCateID: tagcateid, ForumID: forumid}
	_, err = model.GetTagForum(whereMap)
	if err != nil {
		model.AddTagForum(&whereMap)
	}
	return
}

/**
* @des 新增分类
* @param tagCate *TagCate 实例
* @return TagCate, error
 */
func DeleteTagForum(forumid int) (err error) {
	whereMap := map[string]interface{}{"forum_id": forumid}
	err = model.DeleteTagForum(whereMap)
	return
}

/**
* @des 新增分类
* @param tagCate *TagCate 实例
* @return TagCate, error
 */
func GetTagForumsByForumID(forumid int) (tags []model.TagForum, err error) {
	whereMap := &model.TagForum{ForumID: forumid}
	tags, err = model.GetTagForums(whereMap)
	return
}
