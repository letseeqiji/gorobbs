package model

type TagForum struct {
	TagID   int `gorm:"primary_key" json:"tag_id"`   //
	ForumID int `gorm:"primary_key" json:"forum_id"` //
}

/**
* @des 新建一个模型
* @return Tag
 */
func NewTagForum(tagid, forumid int) *TagForum {
	return &TagForum{
		TagID:   tagid,
		ForumID: forumid,
	}
}

/**
* @des 新增tag
* @param tagCate *TagCate 实例
* @return TagCate, error
 */
func AddTagForum(t *TagForum) (tag *TagForum, err error) {
	tag = t
	err = db.Create(tag).Error
	return
}

/**
* @des 新增tag
* @param tagCate *TagCate 实例
* @return TagCate, error
 */
func GetTagForum(whereMap TagForum) (tag *TagForum, err error) {
	err = db.Model(&TagForum{}).Where(whereMap).First(&tag).Error
	return
}

/**
* @des 新增tag
* @param tagCate *TagCate 实例
* @return TagCate, error
 */
func GetTagForums(whereMap *TagForum) (tags []TagForum, err error) {
	err = db.Model(&TagForum{}).Where(whereMap).Find(&tags).Error
	return
}

/**
* @des 删除
* @param tagCate *TagCate 实例
* @return TagCate, error
 */
func DeleteTagForum(whereMap map[string]interface{}) (err error) {
	err = db.Model(&TagForum{}).Where(whereMap).Delete(&TagForum{}).Error
	return
}
