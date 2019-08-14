package model

import "time"

// 缓存表
type Cache struct {
	Model

	K      string    `json:"k"`      //
	V      string    `json:"v"`      //
	Expiry time.Time `json:"expiry"` //
}
