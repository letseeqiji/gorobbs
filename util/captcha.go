package util

import (
	"fmt"

	"github.com/mojocn/base64Captcha"
)

func CodeCaptchaCreate(height, width int) (string, string) {
	//config struct for Character
	//字符,公式,验证码配置
	var configC = base64Captcha.ConfigCharacter{
		Height: height,
		Width:  width,
		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		Mode:               base64Captcha.CaptchaModeNumber,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsUseSimpleFont:    true,
		IsShowHollowLine:   true,
		IsShowNoiseDot:     true,
		IsShowNoiseText:    true,
		IsShowSlimeLine:    true,
		IsShowSineLine:     true,
		CaptchaLen:         4,
	}

	//create a characters captcha.
	//GenerateCaptcha first parameter is empty string,so the package will generate a random uuid for you.
	idKeyC, capC := base64Captcha.GenerateCaptcha("", configC)
	//write to base64 string.
	base64stringC := base64Captcha.CaptchaWriteToBase64Encoding(capC)

	fmt.Println(idKeyC, base64stringC)
	return idKeyC, base64stringC
}

func VerfiyCaptcha(idkey, verifyValue string) bool {
	verifyResult := base64Captcha.VerifyCaptcha(idkey, verifyValue)
	if verifyResult {
		return true
	}
	return false
}
