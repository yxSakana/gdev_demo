package consts

import "time"

type ContentType int

const (
	NovelCt ContentType = iota
	ImageCt
)

const (
	CachePrefix    = "gdd:c:"
	CacheUser      = CachePrefix + "user"
	CacheNovel     = CachePrefix + "novel"
	CacheImage     = CachePrefix + "image"
	NilCacheExpire = 60 * time.Second
)
