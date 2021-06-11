package user

import (
	"gorobbs/model"
)

// 获取用户id
func GetUserByID(uid int) (user model.User, err error) {
	wmap := map[string]interface{}{"id": uid}
	user, err = model.GetUser(wmap)
	return
}

func GetUserBywWechatUnionID(wechatUnionID string) (user model.User, err error) {
	wmap := map[string]interface{}{"wechat_union_id": wechatUnionID}
	user, err = model.GetUser(wmap)
	return
}

func GetUserByEmail(email string) (user model.User, err error) {
	wmap := map[string]interface{}{"email": email}
	user, err = model.GetUser(wmap)
	return
}

func ResetPassword(newPassword string, uid int) (err error) {
	var wmap = make(map[string]interface{})
	wmap["id"] = uid
	err = model.UpdateUser(wmap, map[string]interface{}{"password": newPassword})
	return
}

func UpdateEmailChecked(email string) (err error) {
	var wmap = make(map[string]interface{})
	wmap["email"] = email
	err = model.UpdateUser(wmap, map[string]interface{}{"email_checked": 1})
	return
}

func ResetAvatar(newAvatar string, uid int) (err error) {
	var wmap = make(map[string]interface{})
	wmap["id"] = uid
	err = model.UpdateUser(wmap, map[string]interface{}{"avatar": "/" + newAvatar})
	return
}

func ResetName(newName string, uid int) (err error) {
	var wmap = make(map[string]interface{})
	wmap["id"] = uid
	err = model.UpdateUser(wmap, map[string]interface{}{"username": newName})
	return
}

// 删除用户id
func DelUserByID(uid int) (err error) {
	wmap := map[string]interface{}{"id": uid}
	err = model.DelUser(wmap)
	return
}

func IsAdmin(ugid int) string {
	if ugid > 0 && ugid < 6 {
		return "1"
	}

	return "0"
}

func IsEmailChecked(email string) string {
	user, err := GetUserByEmail(email)
	if err != nil {
		return "0"
	}

	if user.EmailChecked == 1 {
		return "1"
	}

	return "0"
}

//最新会员
func GetNewestTop12Users() (userList []model.User, err error) {
	return model.GetUsers(20, "id desc", 1)
}
