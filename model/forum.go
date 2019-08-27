package model

// 模块表
type Forum struct {
	Model

	Name            string `gorm:"default:''" json:"name"`            //
	Rank            int    `gorm:"default:0" json:"rank"`             //显示，倒序，数字越大越靠前
	ThreadsCnt      int    `gorm:"default:0" json:"threads_cnt"`      //主题数
	TodaypostsCnt   int    `gorm:"default:0" json:"todayposts_cnt"`   //今日发帖，计划任务每日凌晨０点清空为０
	TodaythreadsCnt int    `gorm:"default:0" json:"todaythreads_cnt"` //今日发主题，计划任务每日凌晨０点清空为０
	Brief           string `json:"brief"`                             //版块简介 允许HTML
	Announcement    string `json:"announcement"`                      //版块公告 允许HTML
	Accesson        int    `gorm:"default:0" json:"accesson"`         //是否开启权限控制
	Orderby         int    `gorm:"default:0" json:"orderby"`          //默认列表排序，0: 顶贴时间 last_date， 1: 发帖时间 tid
	Icon            string `gorm:"default:''" json:"icon"`            //板块是否有 icon，存放最后更新时间
	Moduids         string `gorm:"default:''" json:"moduids"`         //每个版块有多个版主，最多10个： 10*12 = 120，删除用户的时候，如果是版主，则调整后再删除。逗号分隔
	SeoTitle        string `gorm:"default:''" json:"seo_title"`       //SEO 标题，如果设置会代替版块名称
	SeoKeywords     string `gorm:"default:''" json:"seo_keywords"`    //
	DigestsNum      int    `gorm:"default:0" json:"digests_num"`      //
}

// 构建forum 模型
func NewForum(name, brief string) *Forum {
	forum := &Forum{}
	return forum
}

// 给定条件 获取指定的forum
func GetForum(maps interface{}) (forum Forum, err error) {
	err = db.Model(&Forum{}).Where(maps).First(&forum).Error
	return
}

// 根据指定的id获取forum
func GetForumByID(id int) (forum Forum, err error) {
	err = db.Model(&Forum{}).Where("id = ?", id).First(&forum).Error
	return
}

// 获取forum的列表【限定排序】
func GetForumsList(orderby string) (forums []Forum, err error) {
	//err = db.Model(&Forum{}).Order("id asc").Find(&forums).Error
	if len(orderby) == 0 {
		orderby = "rank desc"
	}
	err = db.Model(&Forum{}).Order(orderby).Find(&forums).Error
	return
}

// 获取forum的数量
func GetForumTotal(maps interface{}) (count int) {
	db.Model(&Forum{}).Where(maps).Count(&count)

	return
}

// 检测name是否出现过
func ExistForumByName(username string) bool {
	var user User
	db.Model(&User{}).Select("id").Where("username = ?", username).First(&user)
	if user.ID > 0 {
		return true
	}

	return false
}

// 添加一条forum
func AddForum(forumIcon string, forumName string, forumRank int) (forum *Forum, err error) {

	// 入库
	err = db.Create(&Forum{
		Name:  forumName,
		Rank:  forumRank,
		Icon:  forumIcon,
		Brief: "",
	}).Error

	if err != nil {
		return
	}

	return
}

// 增加模块发帖量
func UpdateForumThreadsCnt(id int, newCnt int) error {
	return db.Model(&Forum{}).Where("id = ?", id).Update("threads_cnt", newCnt).Error
}

// 增加模块今日发帖量
func UpdateForumTodaythreadsCnt(id int, newCnt int) error {
	return db.Model(&Forum{}).Where("id = ?", id).Update("todaythreads_cnt", newCnt).Error
}

// 增加模块今日帖子回复量
func UpdateForumTodaypostsCnt(id int, newCnt int) error {
	return db.Model(&Forum{}).Where("id = ?", id).Updates(Forum{TodaypostsCnt: newCnt}).Error
}

type Results struct {
	Totle int
}

// 统计一共有多少个threads = forum中的threanum之和
func SumAllForumThreads() (threadsCount Results, err error) {
	err = db.Model(&Forum{}).Select("sum(threads_cnt) as threadsNum").Scan(&threadsCount).Error
	return
}

// 删除forum
func DelForumByID(id int) (err error) {
	err = db.Unscoped().Where("id = (?)", id).Delete(&Forum{}).Error
	return
}

// 修改--总方法
func UpdateForum(whereMap map[string]interface{}, items map[string]interface{}) (err error) {
	err = db.Model(&Forum{}).Where(whereMap).Update(items).Error
	return
}
