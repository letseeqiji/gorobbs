package model

// 用户角色分组表
type Group struct {
	Model

	Name            string `gorm:"default:''" json:"name"`           //用户组名称
	Creditsfrom     int    `gorm:"default:0" json:"creditsfrom"`     //积分从
	Creditsto       int    `gorm:"default:0" json:"creditsto"`       //积分到
	Allowread       int    `gorm:"default:0" json:"allowread"`       //允许访问
	Allowthread     int    `gorm:"default:0" json:"allowthread"`     //允许发主题
	Allowpost       int    `gorm:"default:0" json:"allowpost"`       //允许回帖
	Allowattach     int    `gorm:"default:0" json:"allowattach"`     //允许上传文件
	Allowdown       int    `gorm:"default:0" json:"allowdown"`       //允许下载文件
	Allowtop        int    `gorm:"default:0" json:"allowtop"`        //允许置顶
	Allowupdate     int    `gorm:"default:0" json:"allowupdate"`     //允许编辑
	Allowdelete     int    `gorm:"default:0" json:"allowdelete"`     //
	Allowmove       int    `gorm:"default:0" json:"allowmove"`       //
	Allowbanuser    int    `gorm:"default:0" json:"allowbanuser"`    //允许禁止用户
	Allowdeleteuser int    `gorm:"default:0" json:"allowdeleteuser"` //
	Allowviewip     int    `gorm:"default:0" json:"allowviewip"`     //允许查看用户敏感信息
}

func GetUserGroupList() (glist []Group, err error) {
	err = db.Model(&Group{}).Find(&glist).Error
	return
}
