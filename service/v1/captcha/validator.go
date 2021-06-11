package captcha

import "github.com/astaxie/beego/validation"

func UserCaptchaValid(v *validation.Validation, cap_key string, captcha string) {
	ValidCaptchaKey(v, cap_key)
	ValidCaptcha(v, captcha)
}

func ValidCaptchaKey(v *validation.Validation, cap_key string) {
	v.Required(cap_key, "cap_key").Message("唯一key不能为空")
}

func ValidCaptcha(v *validation.Validation, captcha string) {
	v.Required(captcha, "captcha").Message("验证码不能为空")
}
 
