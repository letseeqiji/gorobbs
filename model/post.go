package model

// 我的帖子表
type Post struct {
	Model

	ThreadID    int    `gorm:"default:0" json:"thread_id"`     //主题id
	UserID      int    `gorm:"default:0" json:"user_id"`       //
	Isfirst     int    `gorm:"default:0" json:"isfirst"`       //是否为首帖，与 thread.firstpid 呼应
	Userip      string `gorm:"default:''" json:"userip"`       //发帖时用户ip ip2long()
	ImagesNum   int    `gorm:"default:0" json:"images"`        //附件中包含的图片数
	FilesNum    int    `gorm:"default:0" json:"files"`         //附件中包含的文件数
	Doctype     int    `gorm:"default:0" json:"doctype"`       //类型，0: html, 1: txt; 2: markdown; 3: ubb
	QuotePostId int    `gorm:"default:0" json:"quote_post_id"` //引用哪个 pid，可能不存在
	Message     string `gorm:"default:''" json:"message"`      //内容，用户提示的原始数据
	MessageFmt  string `gorm:"default:''" json:"message_fmt"`  //内容，存放的过滤后的html内容，可以定期清理，减肥
	Audited     int    `gorm:"default:1" json:"audited"`       // 通过审核 默认是通过的  有点模块的贴子需要审核才能显示
	LikesCnt    int    `gorm:"default:0" json:"likes_cnt"`     //点赞数
	User        User
	Thread      Thread
	Attach      []Attach
}

// 根据id获取post
func GetPostById(id int) (post Post, err error) {
	err = db.Model(&Post{}).Where("id = ?", id).First(&post).Error
	return
}

// 获取主题的具体内容：isfirst=1的一条
func GetThreadFirstPostByTid(tid int) (post Post, err error) {
	err = db.Model(&Post{}).Where("thread_id = ?", tid).Where("isfirst = ?", 1).First(&post).Error
	return
}

// 获取主题的帖子列表：不要isfirst=1的一条
func GetThreadPostListByTid(tid int, limit int, page int) (post []Post, err error) {
	err = db.Preload("User").Preload("User.Group").Preload("Attach").Model(&Post{}).Where("thread_id = ?", tid).Where("isfirst = ?", 0).Offset((page - 1) * limit).Limit(limit).Find(&post).Error
	return
}

// 添加帖子
func AddPost(post *Post) (*Post, error) {
	// 入库
	err := db.Create(post).Error
	return post, err
}

// 根据id更新post
func UpdatePost(id int, post Post) (upPost Post, err error) {
	err = db.Model(&Post{}).Where("id = ?", id).Updates(post).Error
	upPost, err = GetPostById(id)
	return
}

// 修改附件的数量
func UpdatePostFilesNum(id int, num int) error {
	return db.Model(&Post{}).Where("id = ?", id).Update("files_num", num).Error
}

// 修改点赞的数量
func UpdatePostLikesNum(id int, num int) error {
	return db.Model(&Post{}).Where("id = ?", id).Update("likes_num", num).Error
}

// 统计当前共有帖子的数量：isfirst=0
func CountPostNum() (postNum int, err error) {
	err = db.Model(&Post{}).Where("isfirst = ?", 0).Count(&postNum).Error
	return
}

// 删除
func DelPostsOfThread(tids []string) (err error) {
	err = db.Unscoped().Where("thread_id in (?)", tids).Delete(&Post{}).Error
	return
}
