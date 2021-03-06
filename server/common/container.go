package common

import (
	"zedis/server/hash"
	"zedis/server/list"
	"zedis/server/set"
	"zedis/server/zset"
)

// StringContainer string容器
var StringContainer = make(map[string]string)

// HashContainer hash容器
var HashContainer = make(map[string]*hash.Map)

// ListContainer list容器
var ListContainer = make(map[string]*list.List)

// SetContainer set容器
var SetContainer = make(map[string]*set.Set)

// ZsetContainer zset容器
var ZsetContainer = make(map[string]*zset.Zset)
