package model

// 置顶
type ThreadTop struct {
	ThreadID int `gorm:"primary_key" json:" thread_id "` //
	ForumID  int `json:"forum_id"`                       //查找板块置顶
	Top      int `json:"top"`                            //top: 0 是普通最新贴，> 0 置顶贴
	Thread   Thread
}

// 创建一个置顶
func AddThreadToTop(thread *ThreadTop) (threadTop ThreadTop, err error) {
	err = db.Model(&ThreadTop{}).Create(thread).Error
	return
}

// 修改置顶选项
func UpdateThreadTopTo(threadID int, top int) (err error) {
	err = db.Model(&ThreadTop{}).Where(ThreadTop{ThreadID: threadID}).Updates(ThreadTop{Top: top}).Error
	return
}

// 获得
func GetThreadTopByTid(threadID int) (threadTop ThreadTop, err error) {
	err = db.Model(&ThreadTop{}).Where(ThreadTop{ThreadID: threadID}).First(&threadTop).Error
	return
}

// 删除
func DelThreadTopByTid(threadID int) (threadTop ThreadTop, err error) {
	err = db.Model(&ThreadTop{}).Where(ThreadTop{ThreadID: threadID}).Delete(&threadTop).Error
	return
}

// 获取全站置顶的帖子
func GetTopThreadsWholeWebSite() (threadTop []ThreadTop, err error) {
	err = db.Model(&ThreadTop{}).Preload("Thread").Preload("Thread.User").Where("top = ?", 3).Find(&threadTop).Error
	return
}

// 获取置顶模块下置顶的帖子
func GetTopThreadsForum(forumId int) (threadTop []ThreadTop, err error) {
	err = db.Preload("Thread").Preload("Thread.User").Where("(forum_id = ?) or (forum_id != ? and top = 3) ", forumId, forumId).Find(&threadTop).Error
	return
}

// 删除
func DelthreadTopsOfThread(tids []string) (err error) {
	err = db.Unscoped().Where("thread_id in (?)", tids).Delete(&ThreadTop{}).Error
	return
}
