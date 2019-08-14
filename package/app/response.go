package app

import (
	"gorobbs/package/rcode"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func JsonResponse(c *gin.Context, httpCode, errCode int, data interface{}) {
	c.JSON(httpCode, Response{
		Code: httpCode,
		Msg:  rcode.GetMessage(errCode),
		Data: data,
	})
	return
}
