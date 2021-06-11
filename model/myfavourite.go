package model

// 我的帖子表
type MyFavourite struct {
	Model

	UserID   int `gorm:"default:0" json:"user_id"`   //
	ThreadID int `gorm:"default:0" json:"thread_id"` //
	User     User
	Thread   Thread
}

func GetMyFavouriteList(uid int, page int, limit int, Orderby string) (threads []MyFavourite, err error) {
	if page <= 1 {
		page = 1
	}

	if limit == 0 {
		limit = PAGE_SIZE
	}

	err = db.Preload("User").Preload("Thread").Preload("Thread.Forum").Model(&Mythread{}).Where("user_id = ?", uid).Offset((page - 1) * limit).Limit(limit).Order(Orderby).Find(&threads).Error

	return
}

func AddMyFavourite(userID int, threadID int) (myFavourite *MyFavourite, err error) {
	myFavourite = &MyFavourite{
		UserID:   userID,
		ThreadID: threadID,
	}

	// 入库
	err = db.Create(myFavourite).Error

	return
}

func DelMyFavourite(uid int, threadId int) error {
	return db.Unscoped().Where("user_id = ?", uid).Where("thread_id = ?", threadId).Delete(&MyFavourite{}).Error
}

//检查我的收藏
func CheckFavourite(uid int, tid int) (fav int, err error) {
	err = db.Model(&MyFavourite{}).Where("user_id = ? and thread_id = ?", uid, tid).Count(&fav).Error
	return
}

// 删除
func DelMyFavouritesOfThread(tids []string) (err error) {
	err = db.Unscoped().Where("thread_id in (?)", tids).Delete(&MyFavourite{}).Error
	return
}
