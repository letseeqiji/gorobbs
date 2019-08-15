package regex

func IsEmail(email string) (res bool) {
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

	tocheck := makeRegexp(email, pattern)

	return tocheck.VerifyString()
}
