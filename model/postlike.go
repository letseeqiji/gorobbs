package model

// 我的帖子表
type Postlike struct {
	Model

	UserID int `gorm:"default:0" json:"user_id"` //
	PostID int `gorm:"default:0" json:"post_id"` //
	User   User
	Post   Post
}

func GetMyPostlikeList(uid int, page int, limit int, Orderby string) (postlikes []Postlike, err error) {
	if page <= 1 {
		page = 1
	}

	if limit == 0 {
		limit = PAGE_SIZE
	}

	err = db.Preload("User").Preload("Post").Model(&Mythread{}).Where("user_id = ?", uid).Offset((page - 1) * limit).Limit(limit).Order(Orderby).Find(&postlikes).Error

	return
}

func AddPostlike(userID int, postID int) (postlike *Postlike, err error) {
	postlike = &Postlike{
		UserID: userID,
		PostID: postID,
	}

	// 入库
	err = db.Create(postlike).Error

	return
}

func DelPostlike(uid int, postId int) error {
	return db.Unscoped().Where("user_id = ?", uid).Where("post_id = ?", postId).Delete(&Postlike{}).Error
}

//检查我的收藏
func CheckPostlike(uid int, pid int) (like int, err error) {
	err = db.Model(&Postlike{}).Where("user_id = ? and post_id = ?", uid, pid).Count(&like).Error
	return
}
