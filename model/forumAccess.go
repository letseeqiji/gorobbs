package model

// 模块权限表
type ForumAccess struct {
	Model

	GroupId     int `gorm:"default:0" json:"group_id"`    //
	Allowread   int `gorm:"default:0" json:"allowread"`   //
	Allowthread int `gorm:"default:0" json:"allowthread"` //
	Allowpost   int `gorm:"default:0" json:"allowpost"`   //允许回复
	Allowattach int `gorm:"default:0" json:"allowattach"` //允许上传附件
	Allowdown   int `gorm:"default:0" json:"allowdown"`   //允许下载附件
}
