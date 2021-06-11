package model

type TagCate struct {
	ID           uint   `gorm:"primary_key" json:" - "` //
	Name         string `json:"name"`                   //
	Rank         int    `json:"rank"`                   //
	Enable       int    `json:"enable"`                 //
	DefaultTagID int    `json:"default_tag_id"`         //默认值,如果没有，设为全部
	Isforce      int    `json:"isforce"`                //
	Style        string `json:"style"`                  //
	Comment      string `json:"comment"`                //
	Tags         []Tag
}

/**
* @des 新建一个模型
* @return TagCate
 */
func NewTagCate(name, style, comment string, rank, enable, isforce int) *TagCate {
	return &TagCate{
		Name:         name,
		Rank:         rank,
		Enable:       enable,
		DefaultTagID: 0,
		Isforce:      isforce,
		Style:        style,
		Comment:      comment,
	}
}

/**
* @des 新增分类
* @param tagCate *TagCate 实例
* @return TagCate, error
 */
func AddTagCate(tagCate *TagCate) (tagcate *TagCate, err error) {
	tagcate = tagCate
	err = db.Create(tagCate).Error
	return
}

/**
* @des 获取列表
* @return TagCates, error
 */
func GetTagCateList() (tagcates []TagCate, err error) {
	err = db.Preload("Tags").Find(&tagcates).Error
	return
}

/**
* @des 获取列表
* @return TagCates, error
 */
func GetTagCate(whereMap map[string]interface{}) (tagcate TagCate, err error) {
	err = db.Model(&TagCate{}).Preload("Tags").Where(whereMap).First(&tagcate).Error
	return
}

/**
* @des 修改
* @return TagCates, error
 */
func UpdateTagCate(whereMap map[string]interface{}, items map[string]interface{}) (err error) {
	err = db.Model(&TagCate{}).Where(whereMap).Update(items).Error
	return
}
