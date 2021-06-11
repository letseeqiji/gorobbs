package tag

import "gorobbs/model"

/**
* @des 新增分类
* @param tagCate *TagCate 实例
* @return TagCate, error
 */
func AddTag(cateid int, name, style, comment string, rank, enable int) (newTag *model.Tag, err error) {
	tag := model.NewTag(cateid, name, style, comment, rank, enable)
	newTag, err = model.AddTag(tag)
	return
}

/**
* @des 新增分类
* @param tagCate *TagCate 实例
* @return TagCate, error
 */
func UpdateTag(id, cateid int, name, style, comment string, rank, enable int) (err error) {
	whereMap := map[string]interface{}{"id": id}
	items := map[string]interface{}{"tag_cate_id": cateid, "name": name, "rank": rank, "enable": enable, "style": style, "comment": comment}
	err = model.UpdateTag(whereMap, items)
	return
}

/**
* @des 根据id获取tagcate
* @param id tagcateid
* @return TagCate, error
 */
func GetTagByID(id int) (tag model.Tag, err error) {
	whereMap := map[string]interface{}{"id": id}
	tag, err = model.GetTag(whereMap)
	return
}
