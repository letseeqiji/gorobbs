package model

import "time"

// 任务队列表
type SessionData struct {
	SessionId string    `gorm:"primary_key" json:" - "` //
	LastDate  time.Time `json:"last_date"`              //
	Data      string    `json:"data"`                   //
}
