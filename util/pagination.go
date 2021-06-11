package util

import (
	"gorobbs/package/setting"
	"math"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.ServerSetting.PageSize
	}

	return result
}

func Pagination_tpl(url, text, active string) string {
	g_pagination_tpl := "<li class='page-item{active}'><a href='{url}' class='page-link'>{text}</a></li>"
	g_pagination_tpl = strings.Replace(g_pagination_tpl, "{url}", url, 1)
	g_pagination_tpl = strings.Replace(g_pagination_tpl, "{text}", text, 1)
	g_pagination_tpl = strings.Replace(g_pagination_tpl, "{active}", active, 1)
	return g_pagination_tpl
}

//bootstrap 翻页，命名与 bootstrap 保持一致
// Pagination("?page={page}", 30, 1, 20)
func Pagination(url string, totalnum int, page int, pagesize int) (s string) {

	//fmt.Println(url, "\n", totalnum, "\n", page, "\n", pagesize)

	if pagesize == 0 {
		pagesize = 20
	}
	// totalpage= 100/20=5
	totalpage := math.Ceil(float64(totalnum) / float64(pagesize))
	//fmt.Println("总页数:", totalpage)
	if totalpage < 2 {
		return
	}
	// page = min(5, 2) = 2
	page = int(math.Min(float64(totalpage), float64(page)))
	shownum := 5 // 显示多少个页 * 2

	// start = max(1, 2-5) = 1
	start := int(math.Max(1, float64(page-shownum)))
	// end = max(5, 2+5) = 7
	end := int(math.Min(totalpage, float64(page+shownum)))
	//fmt.Println(start, "-》", end)

	// 不足 $shownum，补全左右两侧
	// right = 2 + 5 - 5 = 2
	right := page + shownum - int(totalpage)
	if right > 0 {
		start -= right
		start = int(math.Max(1, float64(start)))
	}
	// left = 2 - 5 = -3
	left := page - shownum
	if left < 0 {
		end -= left // end = 7+3 = 10
		// end = min(5, 10) = 5
		end = int(math.Min(totalpage, float64(end)))
	}

	// page = 2
	if page != 1 {
		// url = "?page={page}"   -- ?page=1
		url := strings.Replace(url, "{page}", strconv.Itoa(page-1), 1)
		//s = "<li class='page-item'><a href='?page=1' class='page-link'>◀</a></li>"
		s += Pagination_tpl(url, "◀", "")
	}

	// start = 1
	if start > 1 {
		text := "1 "
		if start > 2 {
			text += "..."
		}
		url := strings.Replace(url, "{page}", "1", 1)
		s += Pagination_tpl(url, text, "")
	}

	// for i =1; i < 5; i ++
	for i := start; i <= end; i++ {
		text := ""
		if i == page {
			text += " active"
		}
		// url =
		url := strings.Replace(url, "{page}", strconv.Itoa(i), 1)
		s += Pagination_tpl(url, strconv.Itoa(i), text)
	}

	if end != int(totalpage) {
		text := ""
		if (int(totalpage) - end) > 1 {
			text = "..." + strconv.Itoa(int(totalpage))
		} else {
			text = strconv.Itoa(int(totalpage))
		}
		url := strings.Replace(url, "{page}", strconv.Itoa(int(totalpage)), 1)
		s += Pagination_tpl(url, text, "")
	}

	if page != int(totalpage) {
		url := strings.Replace(url, "{page}", strconv.Itoa(page+1), 1)
		s += Pagination_tpl(url, "▶", "")
	}

	return
}
