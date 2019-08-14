package model

// 任务队列表
type Queue struct {
	ID     int `gorm:"primary_key" json:"id"`    //
	Value  int `gorm:"primary_key" json:"value"` //
	Expiry int `json:"expiry"`                   //过期时间
}
