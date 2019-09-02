package model

type Tag struct {
	ID        int    `gorm:"primary_key" json:"id"` //
	TagCateID int    `json:"tag_cate_id"`           //cate所属分类
	Name      string `json:"name"`                  //
	Rank      int    `json:"rank"`                  //
	Enable    int    `json:"enable"`                //
	Style     string `json:"style"`                 //
	Comment   string `json:"comment"`               //
	TagCate   TagCate
}

/**
* @des 新建一个模型
* @return Tag
 */
func NewTag(cateid int, name, style, comment string, rank, enable int) *Tag {
	return &Tag{
		TagCateID: cateid,
		Name:      name,
		Rank:      rank,
		Enable:    enable,
		Style:     style,
		Comment:   comment,
	}
}

/**
* @des 新增tag
* @param tagCate *TagCate 实例
* @return TagCate, error
 */
func AddTag(t *Tag) (tag *Tag, err error) {
	tag = t
	err = db.Create(tag).Error
	return
}

/**
* @des 获取列表
* @return TagCates, error
 */
func GetTag(whereMap map[string]interface{}) (tag Tag, err error) {
	err = db.Model(&Tag{}).Preload("TagCate").Where(whereMap).First(&tag).Error
	return
}

/**
* @des 修改
* @return TagCates, error
 */
func UpdateTag(whereMap map[string]interface{}, items map[string]interface{}) (err error) {
	err = db.Model(&Tag{}).Where(whereMap).Update(items).Error
	return
}
