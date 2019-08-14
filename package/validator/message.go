package validator

import (
	rcode "gorobbs/package/rcode"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

func VErrorMsg(c *gin.Context, v *validation.Validation, code int) {
	vmsg := make(map[string]interface{})
	for _, err := range v.Errors {
		vmsg[err.Key] = err.Message
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  rcode.GetMessage(code),
		"data": vmsg,
	})
}
