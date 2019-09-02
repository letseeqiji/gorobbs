package model

import "time"

type IpBanned struct {
	IP       string    `json:"ip"`        //
	FromDate time.Time `json:"from_date"` //
	ToDate   time.Time `json:"to_date"`   //
}

/**
* @des 新建一个模型
* @return Tag
 */
func NewIpBanned(ip string, FromDate, ToDate time.Time) *IpBanned {
	return &IpBanned{
		IP:       ip,
		FromDate: FromDate,
		ToDate:   ToDate,
	}
}

/**
* @des 新增tag
* @param tagCate *TagCate 实例
* @return TagCate, error
 */
func AddIpBanned(t *IpBanned) (tag *IpBanned, err error) {
	tag = t
	err = db.Create(tag).Error
	return
}

/**
* @des 新增tag
* @param tagCate *TagCate 实例
* @return TagCate, error
 */
func CheckIpBanned(ip string) (banned bool) {
	res := 0
	db.Model(&IpBanned{}).Where("ip = ?", ip).Count(&res)
	if res != 0 {
		return true
	}
	return false
}
