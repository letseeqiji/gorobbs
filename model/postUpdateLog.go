package model

// 我的帖子表
type PostUpdateLog struct {
	Model

	PostID    	int    `gorm:"default:0" json:"thread_id"`     //主题id
	UserID      int    `gorm:"default:0" json:"user_id"`       //
	Reason     string `gorm:"default:''" json:"reason"`      //内容，用户提示的原始数据
	Message     string `gorm:"default:''" json:"message"`      //内容，用户提示的原始数据
	OldMessage  string `gorm:"default:''" json:"message_fmt"`  //内容，存放的过滤后的html内容，可以定期清理，减肥
	Audited     int    `gorm:"default:1" json:"audited"`       // 通过审核 默认是通过的  有点模块的贴子需要审核才能显示
	User        User
	Post		Post
}

// 添加帖子
func AddPostUpdateLog(postUpdateLog *PostUpdateLog) (*PostUpdateLog, error) {
	// 入库
	err := db.Model(&Thread{}).Create(postUpdateLog).Error

	return postUpdateLog, err
}


