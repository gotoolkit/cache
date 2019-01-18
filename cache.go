package cache

import (
	"sync"

	"github.com/gotoolkit/lru"
)

// cache 封装*lru.Cache
// 添加同步
type cache struct {
	mu         sync.RWMutex
	nbytes     int64 // 键值bytes
	lru        *lru.Cache
	nhit, nget int64
	nevict     int64
}
