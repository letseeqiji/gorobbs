package layout

import (
	"gorobbs/model"
)

func GetForumList() (forums []model.Forum) {
	// 首先列出已经具有的模块
	forums, _ = model.GetForumsList("id asc")
	return
}
