package web

import (
	"github.com/gin-gonic/gin"
	"github.com/huichen/wukong/engine"
	"github.com/huichen/wukong/types"
	"gorobbs/model"
	"gorobbs/package/setting"
	"gorobbs/service/v1/user"
	"log"
	"math"
	"net/http"
	"strconv"
	"unicode/utf8"
)

var (
	// searcher是协程安全的
	searcher = engine.Engine{}
)

func SearchInit()  {
	// 初始化
	searcher.Init(types.EngineInitOptions{SegmenterDictionaries: "static/searcher/data/dictionary.txt"})

	threadCnt, _ := model.CountThreadsNum()
	page := int(math.Ceil(float64(threadCnt) / 20))

	log.Print("总数", threadCnt, page)

	whereMap := &model.Thread{Isclosed: 0}
	order := "created_at desc"
	limit := 20
	// 导入索引：要搜什么就把什么变成索引
	for i := 1; i <= page; i ++ {
		list, _ := model.GetThreads(whereMap, order, limit, i)
		for _,v := range list {
			log.Print(v.Subject)
			AddSearchIndex(uint64(v.ID), v.Subject)
		}
	}
	// 等待索引刷新完毕
	searcher.FlushIndex()
}

func AddSearchIndex(id uint64, subject string)  {
	searcher.IndexDocument(id, types.DocumentIndexData{Content: subject}, false)
}

func Search(c *gin.Context) {
	// 用户是否登录
	islogin := user.IsLogin(c)
	// 获取sessions
	sessions := user.GetSessions(c)
	// 网站描述
	webname := setting.ServerSetting.Sitename
	description := setting.ServerSetting.Sitebrief
	forumname := "搜索"

	var threads []*model.Thread

	key := c.DefaultQuery("key", "")
	defer searcher.Close()
	if utf8.RuneCountInString(key) > 1 {
		output := searcher.Search(types.SearchRequest{Text:key,RankOptions: &types.RankOptions{
			OutputOffset:    0,
			MaxOutputs:      100,
		}})
		// 搜索输出格式见types.SearchResponse结构体
		var ids []string
		for _, doc := range output.Docs {
			ids = append(ids, strconv.Itoa(int(doc.DocId)))
		}
		threads, _ = model.GetThreadsByIDs(ids)
	}

	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"search.html",
		// Pass the data that the page uses
		gin.H{
			"forums":     	forums,
			"islogin":    	islogin,
			"sessions":   	sessions,
			"webname":	webname,
			"description":	description,
			"forumname":	forumname,
			"keyword" : 	key,
			"threads":	threads,
		},
	)
}
