package tag

import "gorobbs/model"

/**
* @des 新增分类
* @param tagCate *TagCate 实例
* @return TagCate, error
 */
func AddTagCate(name, style, comment string, rank, enable, isforce int) (err error) {
	cateCate := model.NewTagCate(name, style, comment, rank, enable, isforce)
	_, err = model.AddTagCate(cateCate)
	return
}

/**
* @des 修改分类
* @param tagCate *TagCate 实例
* @return TagCate, error
 */
func UpdateTagCate(id int, name, style, comment string, rank, enable, isforce, defaultTagid int) (err error) {
	whereMap := map[string]interface{}{"id": id}
	items := map[string]interface{}{"name": name, "rank": rank, "enable": enable, "style": style, "cmooent": comment, "isforce": isforce, "default_tag_id": defaultTagid}
	err = model.UpdateTagCate(whereMap, items)
	return
}

/**
* @des 修改默认值
* @param id tagcateid
* @param tagid tagid
* @return TagCate, error
 */
func UpdateTagCateDefaultTagIDByID(id int, tagId int) (err error) {
	whereMap := map[string]interface{}{"id": id}
	items := map[string]interface{}{"default_tag_id": tagId}
	err = model.UpdateTagCate(whereMap, items)
	return
}

/**
* @des 根据id获取tagcate
* @param id tagcateid
* @return TagCate, error
 */
func GetTagCateByID(id int) (tagCate model.TagCate, err error) {
	whereMap := map[string]interface{}{"id": id}
	tagCate, err = model.GetTagCate(whereMap)
	return
}
