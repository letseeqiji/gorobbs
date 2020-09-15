package v1

import (
	"bytes"
	"html/template"
	"math"
	"net"
	"strconv"
	"time"
)

/**
* @des 模板楼层数字增加1
* @param x string 要转换的字符串）
* @return string
 */
func selfPlus(num int) int {
	return num + 1
}

/**
* @des 避免模板某些字段自动转义实体
* @param x string 要转换的字符串）
* @return string
 */
func unescaped(x string) interface{} {
	return template.HTML(x)
}

/**
* @des 时间转换函数
* @param atime string 要转换的时间戳（秒）
* @return string
 */
func StrTime(atime int64) string {
	var byTime = []int64{365 * 24 * 60 * 60, 24 * 60 * 60, 60 * 60, 60, 1}
	var unit = []string{"年前", "天前", "小时前", "分钟前", "秒钟前"}
	now := time.Now().Unix()
	ct := now - atime
	if ct < 0 {
		return "刚刚"
	}
	var res string
	for i := 0; i < len(byTime); i++ {
		if ct < byTime[i] {
			continue
		}
		var temp = math.Floor(float64(ct / byTime[i]))
		ct = ct % byTime[i]
		if temp > 0 {
			var tempStr string
			tempStr = strconv.FormatFloat(temp, 'f', -1, 64)
			res = MergeString(tempStr, unit[i]) //此处调用了一个我自己封装的字符串拼接的函数（你也可以自己实现）
		}
		break //我想要的形式是精确到最大单位，即："2天前"这种形式，如果想要"2天12小时36分钟48秒前"这种形式，把此处break去掉，然后把字符串拼接调整下即可（别问我怎么调整，这如果都不会我也是无语）
	}
	return res
}

/**
* @des 拼接字符串
* @param args ...string 要被拼接的字符串序列
* @return string
 */
func MergeString(args ...string) string {
	buffer := bytes.Buffer{}
	for i := 0; i < len(args); i++ {
		buffer.WriteString(args[i])
	}
	return buffer.String()
}

/**
* @des 拼接字符串
* @param args ...string 要被拼接的字符串序列
* @return string
 */
func ToMyOrUser(uid int) string {
	return ""
}

/**
* @des 数据自增
* @param args ...string 要被拼接的字符串序列
* @return string
 */
func numPlusPlus(num int) int {
	num++
	return num
}

/**
* @des 把数字ip转化为普通ip
* @param args ...string 要被拼接的字符串序列
* @return string
 */
func Long2IPString(si string) string {
	i, _ := strconv.Atoi(si)
	ip := make(net.IP, net.IPv4len)
	ip[0] = byte(i >> 24)
	ip[1] = byte(i >> 16)
	ip[2] = byte(i >> 8)
	ip[3] = byte(i)

	return ip.String()
}

/**
* @des 把数字ip转化为普通ip
* @param args ...string 要被拼接的字符串序列
* @return string
 */
func SubStr(si string) string {
	slen := 7
	if len([]rune(si)) >= slen {
		si = string([]rune(si)[:slen])
	}
	return si

}
