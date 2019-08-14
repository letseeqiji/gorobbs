package file

import (
	"math/rand"
	"os"
	"strconv"
	"time"
	time_package "gorobbs/package/time"
)

/**
如果返回的错误为nil,说明文件或文件夹存在
如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
如果返回的错误为其它类型,则不确定是否在存在
 */
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreatePath(pathName string) (err error) {
	fileExist, _ := PathExists(pathName)

	if !fileExist {
		err = os.Mkdir(pathName, 0777)
	}

	return
}

func CreatePathInToday(pathName string) (pathInToday string, err error) {
	err = CreatePath(pathName)
	if err != nil {
		return
	}

	today := time_package.TimeFormat("Ymd")
	pathInToday = pathName + "/" + today
	fileExist, _ := PathExists(pathInToday)

	if !fileExist {
		err = os.Mkdir(pathInToday, 0777)
	}

	return
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// makefilename
func MakeFileName(userid string, fileName string) (newFilename string) {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(10000) + 999
	newFilename = userid + "_" + strconv.Itoa(int(time.Now().UnixNano())) + strconv.Itoa(randNum) + GetExt(fileName)
	return newFilename
}
