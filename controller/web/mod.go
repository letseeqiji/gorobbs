package web

import (
	layout_service "gorobbs/service/v1/layout"

	"github.com/gin-gonic/gin"
)

func MoveMod(c *gin.Context) {
	forums := layout_service.GetForumList()
	c.HTML(200, "move_mod.html", gin.H{"forums": forums})
}

func DeleteMod(c *gin.Context) {
	c.HTML(200, "del_mod.html", gin.H{})
}

func TopMod(c *gin.Context) {
	c.HTML(200, "top_mod.html", gin.H{})
}

func CloseMod(c *gin.Context) {
	c.HTML(200, "close_mod.html", gin.H{})
}
