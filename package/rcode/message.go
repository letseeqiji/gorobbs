package rcode

var MessageFlags = map[int]string{
	SUCCESS:                        "操作成功",
	ERROR:                          "系统错误",
	UNPASS:                         "未通过验证",
	BANNED:                         "干坏事儿了吧亲<_>~~，您的访问被拒绝了",
	INVALID_PARAMS:                 "请求参数错误",
	INVALID_CONTENT:                "包含非法内容",
	ERROR_EXIST_TAG:                "已存在该标签名称",
	ERROR_NOT_EXIST_TAG:            "该标签不存在",
	ERROR_NOT_EXIST_ARTICLE:        "该文章不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
	ERROR_EXIST_USER:               "用户已经存在了",
	ERROR_NOT_EXIST_USER:           "用户名或密码错误",
	ERROR_BIND_DATA:                "数据绑定错误",
	ERROR_UNFIND_DATA:              "数据未找到",
	ERROR_IMAGE_BAD_EXT:            "图片格式不正确",
	ERROR_IMAGE_TOO_LARGE:          "图片太大了",
	ERROR_FILE_CREATE_FAIL:         "文件夹创建失败",
	ERROR_FILE_SAVE_FAIL:           "文件保存失败",

	ERROR_SQL_INSERT_FAIL: "数据插入失败",
	ERROR_SQL_DELETE_FAIL: "数据删除失败",
	ERROR_SQL_UPDATE_FAIL: "数据修改失败",
}

func GetMessage(code int) string {
	msg, ok := MessageFlags[code]
	if ok {
		return msg
	}

	return MessageFlags[ERROR]
}
