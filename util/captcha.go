package util

import (
	"fmt"

	"gorobbs/tools/captcha"
)

func CodeCaptchaCreate(height, width int) (string, string) {
	//config struct for Character
	//字符,公式,验证码配置
	var configC = captcha.ConfigCharacter{
		Height: height,
		Width:  width,
		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		Mode:               captcha.CaptchaModeNumber,
		ComplexOfNoiseText: captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  captcha.CaptchaComplexLower,
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
	idKeyC, capC := captcha.GenerateCaptcha("", configC)
	//write to base64 string.
	base64stringC := captcha.CaptchaWriteToBase64Encoding(capC)

	fmt.Println(idKeyC, base64stringC)
	return idKeyC, base64stringC
}

func VerfiyCaptcha(idkey, verifyValue string) bool {
	verifyResult := captcha.VerifyCaptcha(idkey, verifyValue)
	if verifyResult {
		return true
	}
	return false
}
