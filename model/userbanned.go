package model

import "time"

type UserBanned struct {
	UserID   int       `gorm:"primary_key" json:"user_id"` //
	FromDate time.Time `json:"from_date"`                  //
	ToDate   time.Time `json:"to_date"`                    //
}

/**
* @des 新建一个模型
* @return Tag
 */
func NewUserBanned(id int, FromDate, ToDate time.Time) *UserBanned {
	return &UserBanned{
		UserID:   id,
		FromDate: FromDate,
		ToDate:   ToDate,
	}
}

/**
* @des 新增tag
* @param tagCate *TagCate 实例
* @return TagCate, error
 */
func AddUserBanned(t *UserBanned) (tag *UserBanned, err error) {
	tag = t
	err = db.Create(tag).Error
	return
}

/**
* @des 新增tag
* @param tagCate *TagCate 实例
* @return TagCate, error
 */
func CheckUserBanned(id int) (banned bool) {
	res := 0
	db.Model(&UserBanned{}).Where("user_id = ?", id).Count(&res)
	if res != 0 {
		return true
	}
	return false
}
