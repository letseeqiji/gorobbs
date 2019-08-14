package model

// 任务队列表
type ThreadDigest struct {
	ThreadId int `gorm:"primary_key" json:" - "`    //
	ForumId  int `gorm:"default:0" json:"forum_id"` //版块 id
	UserId   int `json:"user_id"`                   //
	Digest   int `json:"digest"`                    //
}
