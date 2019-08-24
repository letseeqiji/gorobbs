package search

import (
	"gorobbs/model"
	"log"
	"math"
	"strconv"

	"github.com/huichen/wukong/engine"
	"github.com/huichen/wukong/types"
)

var (
	searcher    = engine.Engine{}
	forceUpdate = false
)

const (
	MaxOutputs   = 100
	InitSqlLimit = 20
)

// 初始化全文搜索引擎--只要初始化一次即可
func SearchInit() {
	// 初始化
	searcher.Init(types.EngineInitOptions{SegmenterDictionaries: "static/searcher/data/dictionary.txt"})

	threadCnt, _ := model.CountThreadsNum()
	page := int(math.Ceil(float64(threadCnt) / float64(InitSqlLimit)))

	log.Print("总数", threadCnt, page)

	whereMap := &model.Thread{Isclosed: 0}
	order := "created_at desc"

	// 导入索引：要搜什么就把什么变成索引
	for i := 1; i <= page; i++ {
		list, _ := model.GetThreads(whereMap, order, InitSqlLimit, i)
		for _, v := range list {
			log.Print(v.Subject)
			AddSearchIndex(uint64(v.ID), v.Subject)
		}
	}
	// 等待索引刷新完毕
	searcher.FlushIndex()
}

// 添加索引
func AddSearchIndex(id uint64, subject string) {
	searcher.IndexDocument(id, types.DocumentIndexData{Content: subject}, forceUpdate)
}

// 删除索引
func RemoveSearchIndex(id uint64) {
	searcher.RemoveDocument(id, forceUpdate)
}

// 搜索关键字
func Search(key string, page int) (output types.SearchResponse) {
	defer searcher.Close()

	output = searcher.Search(types.SearchRequest{Text: key, RankOptions: &types.RankOptions{
		OutputOffset: MaxOutputs * (page - 1),
		MaxOutputs:   MaxOutputs,
	}})

	return
}

// 获取搜索得到的id，这里是由threadid 组成的切片  因为gorm要用这个作为条件
func OutPutIds(output types.SearchResponse) (ids []string) {
	for _, doc := range output.Docs {
		ids = append(ids, strconv.Itoa(int(doc.DocId)))
	}
	return
}
