package app

import (
	"gorobbs/package/rcode"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response setting gin.JSON
func JsonResponse(c *gin.Context, httpCode, errCode int, data interface{}) {
	c.JSON(httpCode, gin.H{
		"code":    errCode,
		"message": rcode.GetMessage(errCode),
		"data":    data,
	})
	return
}

func JsonOkResponse(c *gin.Context, errCode int, data interface{}) {
	JsonResponse(c, http.StatusOK, errCode, data)
}

func JsonErrResponse(c *gin.Context, errCode int) {
	JsonResponse(c, http.StatusOK, errCode, nil)
}

func HtmlResponse(c *gin.Context, httpCode int, htmlfile string, errCode int, data gin.H) {
	c.HTML(httpCode, htmlfile, data)

	return
}
