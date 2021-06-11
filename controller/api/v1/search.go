package v1

import (
	"github.com/huichen/wukong/engine"
)

var (
	// searcher是协程安全的
	searcher = engine.Engine{}
)
