package model

import "time"

// 任务队列表
type Thread struct {
	Model

	ForumID      int       `gorm:"default:0" json:"forum_id"`      //版块 id
	Top          int       `gorm:"default:0" json:"top"`           //置顶级别: 0: 普通主题, 1-3 置顶的顺序
	UserID       int       `gorm:"default:0" json:"user_id"`       //
	Userip       string    `gorm:"default:''" json:"userip"`       //发帖时用户ip ip2long()，主要用来清理
	Subject      string    `gorm:"default:''" json:"subject"`      // 主题
	LastDate     time.Time `json:"last_date"`                      //最后回复时间
	ViewsCnt     int       `gorm:"default:0" json:"views_cnt"`     //查看次数, 剥离出去，单独的服务，避免 cache 失效
	PostsCnt     int       `gorm:"default:0" json:"posts_cnt"`     //回帖数
	FavouriteCnt int       `gorm:"default:0" json:"favourite_cnt"` //被收藏数
	ImagesNum    int       `gorm:"default:0" json:"images_num"`    //附件中包含的图片数
	FilesNum     int       `gorm:"default:0" json:"files_num"`     //附件中包含的文件数
	ModsCnt      int       `gorm:"default:0" json:"mods_cnt"`      //预留：版主操作次数，如果 > 0, 则查询 modlog，显示斑竹的评分
	Isclosed     int       `gorm:"default:0" json:"isclosed"`      //预留：是否关闭，关闭以后不能再回帖、编辑
	FirstPostID  int       `gorm:"default:0" json:"first_post_id"` //首贴 pid
	LastUserID   int       `gorm:"default:0" json:"last_user_id"`  //最近参与的 uid
	LastPostID   int       `gorm:"default:0" json:"last_post_id"`  //最后回复的 pid
	Digest       int       `gorm:"default:0" json:"digest"`        //
	Audited      int       `gorm:"default:1" json:"audited"`       // 通过审核 默认是通过的  有点模块的贴子需要审核才能显示
	User         User      `json:"user"`
	Forum        Forum
	LastUser     User
	Attach       []Attach
}

func GetThreads(whereMap interface{}, order string, limit int, page int) (thread []Thread, err error) {
	err = db.Model(&Thread{}).Preload("User").Where(whereMap).Order(order).Offset((page - 1) * limit).Limit(limit).Find(&thread).Error
	return
}

// 首页用获取帖子列表
func GetThreadList(page int) (threads []Thread, err error) {
	if page <= 1 {
		page = 1
	}

	err = db.Preload("User").Model(&Thread{}).Where("isclosed = ?", 0).Where("top != ?", 3).Order("created_at desc").Offset((page - 1) * PAGE_SIZE).Limit(PAGE_SIZE).Find(&threads).Error

	return
}

// 获取帖子的总数
func GetThreadTotleCount() (totle int, err error) {
	err = db.Model(&Thread{}).Where("isclosed = ?", 0).Where("top != ?", 3).Count(&totle).Error
	return
}

// 获取未置顶的帖子总数
func GetThreadTotal(maps interface{}) (count int) {
	db.Model(&Thread{}).Where("isclosed = ?", 0).Where("top = ?", 0).Where(maps).Count(&count)

	return
}

// 获取分类下所有的thread:[id,subject，查看量，回帖量，创建时间]， user[id, avatar, username]
func GetThreadListByForumID(forumID int, page int, limit int, orderby string) (threads []Thread, err error) {
	//err = db.Model(&Forum{}).Order("id asc").Find(&forums).Error
	if len(orderby) == 0 {
		orderby = "rank desc"
	}
	err = db.Preload("User").Model(&Thread{}).Where("forum_id = ?", forumID).Where("top = ?", 0).Where("isclosed = ?", 0).Offset((page - 1) * limit).Limit(limit).Order(orderby).Find(&threads).Error

	//err = db.Preload("User").Select("user.id, user.username, user.avatar").Model(&Thread{}).Select("id, subject, views_cnt, posts_cnt, created_at").Where("forum_id = ?", forumID).Where("top = ?", 0).Where(Thread{Isclosed: 0}).Offset((page - 1) * limit).Limit(limit).Order(orderby).Find(&threads).Error
	//err = db.Table("bbs_thread as thread").Joins("join bbs_user as user on thread.user_id = user.id").Select("thread.id, thread.subject, user.id, user.username").Where("thread.forum_id = ?", forumID).Scan(&threads).Error
	//err = db.Raw("SELECT thread.id, thread.subject, user.id, user.avatar, user.username FROM bbs_thread as thread, bbs_user as user WHERE thread.user_id = user.id and thread.forum_id = ?", forumID).Scan(&threads).Error

	return
}

// 获取指定分类下所有的帖子的id--目前删除用
func GetThreadIDSByForumID(forumID int) (threads []Thread, err error) {
	err = db.Model(&Thread{}).Select("id").Where("forum_id = ?", forumID).Find(&threads).Error
	return
}

// 根据id获得thread
func GetThreadById(id int) (thread Thread, err error) {
	err = db.Preload("Forum").Preload("User").Where("id = ?", id).Model(&Thread{}).First(&thread).Error
	return
}

// 添加帖子
func AddThread(thread *Thread) (*Thread, error) {
	err := db.Model(&Thread{}).Create(thread).Error
	return thread, err
}

// 修改
func UpdateThread(id int, thread Thread) (upthread Thread, err error) {
	err = db.Model(&Thread{}).Where("id = ?", id).Updates(thread).Error
	upthread, err = GetThreadById(id)
	return
}
func UpdateThreadPro(id int, items map[string]interface{}) (upthread Thread, err error) {
	err = db.Model(&Thread{}).Where("id = ?", id).Updates(items).Error
	upthread, err = GetThreadById(id)
	return
}

// 删除
func DelThread(ids []string) (err error) {
	err = db.Unscoped().Where("id in (?)", ids).Delete(&Thread{}).Error
	return
}

// 增加阅读量
func UpdateThreadViewsCnt(id int) error {
	thread, _ := GetThreadById(id)
	return db.Model(&Thread{}).Where("id = ?", id).Update("views_cnt", thread.ViewsCnt+1).Error
}

// 修改置顶级别
func UpdateThreadTop(id int, top int) error {
	return db.Model(&Thread{}).Where("id = ?", id).Update("top", top).Error
}

// 修改附件的数量
func UpdateThreadFilesNum(id int, num int) error {
	return db.Model(&Thread{}).Where("id = ?", id).Update("files_num", num).Error
}

// 修改附件的数量
func UpdateThreadFavouriteCnt(id int, num int) error {
	return db.Model(&Thread{}).Where("id = ?", id).Update("favourite_cnt", num).Error
}

// 修改帖子到新模块
func UpdateThreadForum(ids string, nfid int) error {
	return db.Model(&Thread{}).Where("id in (?)", ids).Update("forum_id", nfid).Error
}

// 统计一共共有多少thread
func CountThreadsNum() (threadsNum int, err error) {
	err = db.Model(&Thread{}).Count(&threadsNum).Error
	return
}

// 根据id列表获取一组thread
func GetThreadsByIDs(ids []string) (threads []*Thread, err error) {
	err = db.Model(&Thread{}).Preload("User"). /*.Select("id, subject, views_cnt, posts_cnt, user")*/ Where("id in (?)", ids).Where("isclosed = ?", 0).Order("created_at desc").Find(&threads).Error
	return
}
