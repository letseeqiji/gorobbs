package regex

import "regexp"

type Regex struct {
	Str string
	Pattern string
}

/*func VerifyString(str string, pattern string) bool {
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(str)
}
*/
func makeRegexp(str string, pattern string) *Regex {
	return &Regex{
		Str:str,
		Pattern:pattern,
	}
}

func (this *Regex)VerifyString() bool {
	reg := regexp.MustCompile(this.Pattern)
	return reg.MatchString(this.Str)
}
