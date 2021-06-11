package model

// 附件表
type Attach struct {
	Model

	ThreadID     int    `gorm:"default:0" json:"thread_id"`     //主题id
	PostID       int    `gorm:"default:0" json:"post_id"`       //帖子id
	UserID       int    `gorm:"default:0" json:"user_id"`       //用户id
	Filesize     int    `gorm:"default:0" json:"filesize"`      //文件尺寸，单位字节
	Width        int    `gorm:"default:0" json:"width"`         //width > 0 则为图片
	Height       int    `gorm:"default:0" json:"height"`        //
	Filename     string `gorm:"default:''" json:"filename"`     //文件名称，会过滤，并且截断，保存后的文件名，不包含URL前缀 upload_url
	Orgfilename  string `gorm:"default:''" json:"orgfilename"`  //上传的原文件名
	Filetype     string `gorm:"default:''" json:"filetype"`     //image/txt/zip，小图标显示
	Comment      string `gorm:"default:''" json:"comment"`      //文件注释 方便于搜索
	DownloadsNum int    `gorm:"default:0" json:"downloads_num"` //下载次数
	CreditsNum   int    `gorm:"default:0" json:"credits_num"`   //需要的积分
	GoldsNum     int    `gorm:"default:0" json:"golds_num"`     //需要的金币
	RmbsNum      int    `gorm:"default:0" json:"rmbs_num"`      //需要的人民币
	Isimage      int    `gorm:"default:0" json:"isimage"`       //是否为图片
}

// 添加一个附件
func AddAttach(attach *Attach) (*Attach, error) {
	err := db.Model(&Attach{}).Create(attach).Error
	return attach, err
}

// 获取指定post的附件列表
func GetAttachsByPostId(postId int) (attachs []Attach, err error) {
	err = db.Model(&Attach{}).Where("post_id = ?", postId).Find(&attachs).Error
	return
}

// 删除某个附件【应该删除具体的文件】
func DelAttach(id int) error {
	return db.Model(&Attach{}).Where("id = ?", id).Unscoped().Delete(&Attach{}).Error
}

// 统计当前网站共有附件的数量
func CountAttachsNum() (accachsNum int, err error) {
	err = db.Model(&Attach{}).Count(&accachsNum).Error
	return
}

// 删除
func DelAttachsOfThread(tids []string) (err error) {
	err = db.Unscoped().Where("thread_id in (?)", tids).Delete(&Attach{}).Error
	return
}
