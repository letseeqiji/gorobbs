package model

type TagThread struct {
	TagID    int `gorm:"primary_key" json:"tag_id"`    //
	ThreadID int `gorm:"primary_key" json:"thread_id"` //
	ForumID  int `gorm:"primary_key" json:"forum_id"`  //
}

/**
* @des 新建一个模型
* @return Tag
 */
func NewTagThread(tagid, threadid, forumid int) *TagThread {
	return &TagThread{
		TagID:    tagid,
		ThreadID: threadid,
		ForumID:  forumid,
	}
}

/**
* @des 新增tag
* @param tagCate *TagCate 实例
* @return TagCate, error
 */
func AddTagThread(t *TagThread) (tagthread *TagThread, err error) {
	tagthread = t
	err = db.Create(tagthread).Error
	return
}

/**
* @des 赛选
* @param tagCate *TagCate 实例
* @return TagCate, error
 */
func GetTagThreadIDByForumIDWithTagID(id, tagid int) (threadids []TagThread, err error) {
	err = db.Model(&TagThread{}).Select("thread_id").Where("forum_id = ?", id).Where("tag_id = ?", tagid).Find(&threadids).Error
	return
}
