package model

// 缓存表
type Modlog struct {
	Model

	UserId     int    `gorm:"default:0" json:"user_id"`     //
	ThreadId   int    `gorm:"default:0" json:"thread_id"`   //
	PostId     int    `gorm:"default:0" json:"post_id"`     //
	Subject    string `gorm:"default:''" json:"subject"`    //
	Comment    string `gorm:"default:''" json:"comment"`    //
	Rmbs       int    `gorm:"default:0" json:"rmbs"`        //
	CreateDate int    `gorm:"default:0" json:"create_date"` //
	Action     string `gorm:"default:''" json:"action"`     //
}
