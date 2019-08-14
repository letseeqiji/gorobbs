package time

import "time"

func TimeFormat(template string) (formatTime string) {
	switch template {
	case "Ymd":
		formatTime = time.Now().Format("20060102")
	case "Y/m/d":
		formatTime = time.Now().Format("2006/01/02")
	case "Y-m-d":
		formatTime = time.Now().Format("2006-01-02")
	case "H:i:s":
		formatTime = time.Now().Format("15:04:05")
	case "Ymd H:i:s":
		formatTime = time.Now().Format("20060102 15:04:05")
	case "Y-m-d  H:i:s":
		formatTime = time.Now().Format("2006-01-02  15:04:05")
	default:
		formatTime = time.Now().Format("2006/01/02  15:04:05")
	}
	return
}