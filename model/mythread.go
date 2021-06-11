package model

// 我的帖子表
type Mythread struct {
	Model

	UserID   int `gorm:"default:0" json:"user_id"`   //
	ThreadID int `gorm:"default:0" json:"thread_id"` //
	User     User
	Thread   Thread
}

func GetMyThreadList(uid int, page int, limit int, Orderby string) (threads []Mythread, err error) {
	if page <= 1 {
		page = 1
	}

	if limit == 0 {
		limit = 20
	}

	err = db.Preload("User").Preload("Thread").Model(&Mythread{}).Where("user_id = ?", uid).Offset((page - 1) * 20).Limit(200).Order(Orderby).Find(&threads).Error

	return
}

func AddMyThread(userID int, threadID int) (myThread *Mythread, err error) {
	myThread = &Mythread{
		UserID:   userID,
		ThreadID: threadID,
	}

	// 入库
	err = db.Create(myThread).Error

	return
}

// 删除
func DelMyThreadsOfThread(tids []string) (err error) {
	err = db.Unscoped().Where("thread_id in (?)", tids).Delete(&Mythread{}).Error
	return
}
