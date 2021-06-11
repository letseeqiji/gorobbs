package admin

import (
	"gorobbs/model"
	package_online "gorobbs/package/online"
	"gorobbs/service/v1/user"

	"github.com/gin-gonic/gin"
)

func AdminIndex(c *gin.Context) {
	sessions := user.GetSessions(c)

	//主题数：2
	threadsNum, _ := model.CountThreadsNum()
	//帖子数：8
	postsNum, _ := model.CountPostNum()
	//用户数：2
	usersNum, _ := model.CountUserNum()
	//附件总数：0
	attachsNum, _ := model.CountAttachsNum()
	/*磁盘剩余空间：29.15G
	在线人数：1*/
	onlinePeoples, _ := package_online.DbSize()

	/*服务器信息 [PHPINFO]
	操作系统：WINNT
	Web Server：Apache/2.4.39 (Win64) OpenSSL/1.0.2r PHP/7.1.29
	PHP：7.1.29
	数据库：pdo_mysql (10.1.39-MariaDB)
	最大 POST 数据大小：8M
	最大文件上传大小：2M
	允许开启远程 URL：是
	安全模式（safe_mode）：否
	最长执行时间：30
	内存上限：128M
	客户端 IP：127.0.0.1
	服务端 IP：127.0.0.1*/

	/*开发团队信息
	Official Site: http://www.xiuno.com/
	Dev Team: axiuno
	Thanks For:*/

	c.HTML(
		200,
		"aindex.html",
		gin.H{
			"sessions": sessions,
			"counts": map[string]interface{}{
				"threads": threadsNum,
				"posts":   postsNum,
				"users":   usersNum,
				"attachs": attachsNum,
				"onlines": onlinePeoples,
			},
		})
}
