package user

import (
	"regexp"

	"github.com/astaxie/beego/validation"
)

//获取用户 根据用户名
func AddUserValid(v *validation.Validation, username string, password string, email string) {
	ValidName(v, username)
	ValidPassword(v, password)
	// ValidPhone(v, phone)
	ValidEmail(v, email)
}

func LoginValidWithEmail(v *validation.Validation, email string, password string) {
	ValidPassword(v, password)
	ValidEmail(v, email)
}

func LoginValidWithName(v *validation.Validation, name string, password string) {
	ValidPassword(v, password)
	ValidNameRequired(v, name)
}

func ValidName(v *validation.Validation, username string) {
	pass, _ := regexp.MatchString("[a-zA-Z0-9]{3,16}", username)
	if !pass {
		// 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
		v.SetError("username", "名称只能是3-16位字母数字组合")
	}
}

func ValidNameRequired(v *validation.Validation, username string) {
	v.Required(username, "username").Message("密码不能为空")
}

func ValidPassword(v *validation.Validation, password string) {
	v.Required(password, "password").Message("密码不能为空")
}

func ValidPhone(v *validation.Validation, phone string) {
	v.Mobile(phone, "phone").Message("手机不能为空或格式不正确")
}

func ValidEmail(v *validation.Validation, email string) {
	v.Email(email, "email").Message("邮箱不能为空或格式不正确")
}
