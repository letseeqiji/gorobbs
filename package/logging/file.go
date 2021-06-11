package logging

import (
	"fmt"
	"gorobbs/package/setting"
	"time"
)

// getLogFilePath get the log file save path
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.ServerSetting.RuntimeRootPath, setting.ServerSetting.LogSavePath)
}

// getLogFileName get the save name of the log file
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.ServerSetting.LogSaveName,
		time.Now().Format(setting.ServerSetting.TimeFormat),
		setting.ServerSetting.LogFileExt,
	)
}
