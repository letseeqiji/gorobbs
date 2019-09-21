package model

import (
	"gorobbs/util"
	"time"
)

/**
* @des 用户模型
 */
type User struct {
	Model

	Username      string    `gorm:"default:''" json:"username"`        //用户名
	Realname      string    `gorm:"default:''" json:"realname"`        //用户实名
	GroupID       int       `gorm:"default:0" json:"group_id"`         //用户组编号
	Email         string    `gorm:"default:''" json:"email"`           //邮箱
	EmailChecked  int       `gorm:"default:0" json:"email_checked"`    //邮箱验证过
	Password      string    `gorm:"default:''" json:"password"`        //密码
	PasswordSms   string    `gorm:"default:''" json:"password_sms"`    //密码
	Phone         string    `gorm:"default:''" json:"phone"`           //手机号
	IdNumber      string    `gorm:"default:''" json:"id_number"`       //用户名
	Qq            string    `gorm:"default:''" json:"qq"`              //QQ
	WechatUnionID string    `gorm:"default:''" json:"wechat_union_id"` //微信
	ThreadsCnt    int       `gorm:"default:0" json:"threads_cnt"`      //发帖数
	PostsCnt      int       `gorm:"default:0" json:"posts_cnt"`        //回帖数
	FavouriteCnt  int       `gorm:"default:0" json:"favourite_cnt"`    //收藏帖子数
	CreditsNum    int       `gorm:"default:0" json:"credits_num"`      //积分
	GoldsNum      int       `gorm:"default:0" json:"golds_num"`        //金币
	RmbsNum       int       `gorm:"default:0" json:"rmbs_num"`         //人民币
	CreateIp      string    `gorm:"default:''" json:"create_ip"`       //创建时IP
	LoginIp       string    `gorm:"default:''" json:"login_ip"`        //登录时IP
	LoginDate     time.Time `json:"login_date"`                        //登录时间
	LoginsCnt     int       `gorm:"default:0" json:"logins_cnt"`       //登录次数
	Avatar        string    `gorm:"default:''" json:"avatar"`          //用户最后更新图像时间
	DigestsCnt    int       `gorm:"default:0" json:"digests_cnt"`      //精华数
	State         int       `json:"state"`
	Group         Group
}

/**
* @des 根据条件获取一个用户
* @param maps interface{} 条件
* @return User， error
 */
func GetUser(maps interface{}) (user User, err error) {
	err = db.Preload("Group").Model(&User{}).Where(maps).First(&user).Error

	return
}

/**
* @des 根据条件获取一个用户
* @param maps interface{} 条件
* @return User， error
 */
func GetUserByID(id int) (user User, err error) {
	err = db.Preload("Group").Model(&User{}).Where("id = ?", id).First(&user).Error

	return
}

/**
* @des 根据条件获取用户列表
* @param num int 数量
* @param order string 排序
* @param maps where条件
* @return []User， error
 */
func GetUsers(num int, order string, maps interface{}) (user []User, err error) {
	err = db.Preload("Group").Model(&User{}).Order(order).Limit(num).Find(&user).Error

	return
}

/**
* @des 获取指定条件的用户数
* @param maps where条件
* @return count int 总数
 */
func GetUserTotal(maps interface{}) (count int) {
	db.Model(&User{}).Where(maps).Count(&count)

	return
}

/**
* @des 判断用户名是否使用
* @param username string 用户名
* @return bool 是否
 */
func ExistUserByName(username string) bool {
	var user User
	db.Model(&User{}).Select("id").Where("username = ?", username).First(&user)
	if user.ID > 0 {
		return true
	}

	return false
}

/**
* @des 判断邮箱是否使用
* @param email string 邮箱
* @return bool 是否
 */
func ExistUserByEmail(email string) bool {
	var user User
	db.Model(&User{}).Select("id").Where("email = ?", email).First(&user)
	if user.ID > 0 {
		return true
	}

	return false
}

/**
* @des 判断手机号是否使用
* @param phone string 手机号
* @return bool 是否
 */
func ExistUserByPhone(phone string) bool {
	var user User
	db.Model(&User{}).Select("id").Where("phone = ?", phone).First(&user)
	if user.ID > 0 {
		return true
	}

	return false
}

/**
* @des 新增用户
* @param username string 用户名
* @param password string 密码
* @param email string 邮箱
* @param ip string 注册ip
* @return User, error
 */
func AddUser(username, password, email, ip string) (user *User, err error) {
	// 加密密码
	password, err = util.BcryptString(password)
	if err != nil {
		return
	}

	user = &User{
		Username:  username,
		Password:  password,
		Email:     email,
		CreateIp:  ip,
		LoginDate: time.Now(),
	}

	// 入库
	err = db.Create(user).Error

	return
}

/**
* @des 新增用户
* @param username string 用户名
* @param password string 密码
* @param email string 邮箱
* @param ip string 注册ip
* @return User, error
 */
func AddUserPro(userinfo *User) (user *User, err error) {
	// 加密密码
	userinfo.Password, err = util.BcryptString(userinfo.Password)
	if err != nil {
		return
	}

	// 入库
	user = userinfo
	err = db.Create(userinfo).Error

	return
}

/**
* @des 新增用户
* @param whereMaps 条件
* @param updateItems 修改项
* @return error
 */
func UpdateUser(whereMaps interface{}, updateItems map[string]interface{}) (err error) {
	err = db.Model(&User{}).Where(whereMaps).Updates(updateItems).Error
	return
}

/**
* @des 删除用户
* @param whereMaps 条件
* @return error
 */
func DelUser(whereMaps interface{}) (err error) {
	err = db.Unscoped().Where(whereMaps).Delete(&User{}).Error
	return
}

/**
* @des 更新用户发帖数
* @param whereMaps 条件
* @param updateItems 修改项
* @return error
 */
func UpdateUserThreadsCnt(id int, newThreadsCnt int) (err error) {
	err = db.Model(&User{}).Where("id = ?", id).Update("threads_cnt", newThreadsCnt).Error
	return
}

/**
* @des 更新用户回帖数
* @param whereMaps 条件
* @param updateItems 修改项
* @return error
 */
func UpdateUserPostsCnt(id int, newPostsCnt int) (err error) {
	err = db.Model(&User{}).Where("id = ?", id).Update("posts_cnt", newPostsCnt).Error
	return
}

/**
* @des 更新用户积分数
* @param whereMaps 条件
* @param updateItems 修改项
* @return error
 */
func UpdateUserCreditsNum(id int, newCreditsNum int) (err error) {
	err = db.Model(&User{}).Where("id = ?", id).Update("credits_num", newCreditsNum).Error
	return
}

/**
* @des 更新用户积分数
* @param whereMaps 条件
* @param updateItems 修改项
* @return error
 */
func CountUserNum() (usersNum int, err error) {
	err = db.Model(&User{}).Count(&usersNum).Error
	return
}
