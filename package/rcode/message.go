package rcode

var MessageFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	UNPASS:                         "未通过验证",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_EXIST_TAG:                "已存在该标签名称",
	ERROR_NOT_EXIST_TAG:            "该标签不存在",
	ERROR_NOT_EXIST_ARTICLE:        "该文章不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
	ERROR_EXIST_USER:               "用户已经存在了",
	ERROR_NOT_EXIST_USER:           "用户名或密码错误",
}

func GetMessage(code int) string {
	msg, ok := MessageFlags[code]
	if ok {
		return msg
	}

	return MessageFlags[ERROR]
}
