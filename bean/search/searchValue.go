package search

import (
	"json/common"
	simplejson "json/third-pkg/go-simplejson"
)

// search key
// 查询某一个key
func SearchContain(data []byte, key string) bool {
	if data == nil {
		return false
	}
	j, err := simplejson.NewJson(data)
	if common.CheckError(err) {
		return false
	}
	value, err := j.Get(key).Bool()
	if common.CheckError(err) {
		return false
	}
	return value
}
