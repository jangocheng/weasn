package cache

import (
	"github.com/prodbox/weasn/kernel/cache/internal"
	"os"
	"path"
)

// 缓存接口
type Cache interface {
	Start() error
	// Set 设置缓存(键、值。缓存时间)
	Set(string, interface{}, int64) error
	// Get 获取缓存
	Get(string) interface{}
	// Delete 删除缓存
	Delete(string) error
	// Has 缓存是否存在
	IsExist(string) bool
	// Flush 清空缓存
	Flush() error
}

// 默认缓存目录
var DefaultDir = path.Join(os.TempDir(), "weasn/caches")

// 默认缓存驱动
var DefaultCache = internal.NewFileCacher(DefaultDir)
