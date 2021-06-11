package v1

import (
	"gorobbs/package/app"
	"gorobbs/package/rcode"
	"gorobbs/package/validator"
	captcha_service "gorobbs/service/v1/captcha"
	"gorobbs/util"
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

func GetCapacha(c *gin.Context) {
	height, _ := strconv.Atoi(c.DefaultQuery("height", "60"))
	width, _ := strconv.Atoi(c.DefaultQuery("width", "200"))
	code := rcode.SUCCESS
	data := make(map[string]interface{})

	cap_key, captcha_base64 := util.CodeCaptchaCreate(height, width)
	data["cap_key"] = cap_key
	data["captcha_base64"] = captcha_base64

	app.JsonOkResponse(c, code, data)
}

func VerfiyCaptcha(c *gin.Context) {
	cap_key := c.DefaultPostForm("cap_key", "")
	captcha := c.DefaultPostForm("captcha", "")
	code := rcode.SUCCESS
	data := make(map[string]interface{})

	valid := &validation.Validation{}
	captcha_service.UserCaptchaValid(valid, cap_key, captcha)
	if valid.HasErrors() {
		code = rcode.INVALID_PARAMS
		validator.VErrorMsg(c, valid, code)
		return
	}

	pass := util.VerfiyCaptcha(cap_key, captcha)
	if !pass {
		code = rcode.UNPASS
	}

	app.JsonOkResponse(c, code, data)
}
