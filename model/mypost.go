package model

// 我的帖子表
type Mypost struct {
	Model

	UserID   int `gorm:"primary_key;default:0" json:"user_id"` //
	ThreadID int `gorm:"default:0" json:"thread_id"`           //
	PostID   int `gorm:"primary_key;default:0" json:"post_id"` //
	User     User
	Thread   Thread
	Post     Post
}

func GetMyPostList(uid int, page int, limit int, Orderby string) (posts []Mypost, err error) {
	if page <= 1 {
		page = 1
	}

	if limit == 0 {
		limit = 20
	}

	err = db.Preload("Thread").Preload("Thread.User").Preload("Thread.User.Group").Preload("Post").Model(&Mypost{}).Where("user_id = ?", uid).Offset((page - 1) * 20).Limit(200).Order(Orderby).Find(&posts).Error

	return
}

func AddMyPost(userID int, threadID int, postID int) (myPost *Mypost, err error) {
	myPost = &Mypost{
		UserID:   userID,
		ThreadID: threadID,
		PostID:   postID,
	}

	// 入库
	err = db.Create(myPost).Error

	return
}

// 删除
func DelMyPostsOfThread(tids []string) (err error) {
	err = db.Unscoped().Where("thread_id in (?)", tids).Delete(&Mypost{}).Error
	return
}
