package regex

func IsMobile(mobile string) (res bool) {
	pattern := `^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$`

	tocheck := makeRegexp(mobile, pattern)

	return tocheck.VerifyString()
}

