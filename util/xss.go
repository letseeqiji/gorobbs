package util

import (
	"github.com/microcosm-cc/bluemonday"
)

func XssPolice(html string) string {
	p := bluemonday.UGCPolicy()

	return p.Sanitize(html)
}
