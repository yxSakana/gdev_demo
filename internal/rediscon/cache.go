package rediscon

import (
	"fmt"
	"github.com/yxSakana/gdev_demo/internal/consts"
)

type Cacheable interface {
	GetCacheKey(id uint64) string
	ToCacheMap() map[string]interface{}
	RefreshFromCacheMap(cacheRet map[string]interface{})
	SaveToCache()
}

func GetCacheKey(id uint64, ct consts.ContentType) (key string) {
	switch ct {
	case consts.NovelCt:
		key = fmt.Sprintf("%s:%d", consts.CacheNovel, id)
	case consts.ImageCt:
		key = fmt.Sprintf("%s:%d", consts.CacheImage, id)
	}
	return
}
