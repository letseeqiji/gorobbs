package string

import "strconv"

func A2i(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

func I2A(i int) string {
	a := strconv.Itoa(i)
	return a
}
