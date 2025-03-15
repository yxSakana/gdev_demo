package consts

type ContentType int

const (
	NovelCt ContentType = iota
	ImageCt
)

const (
	CachePrefix = "gdd:c:"
	CacheUser   = CachePrefix + "user"
	CacheNovel  = CachePrefix + "novel"
	CacheImage  = CachePrefix + "image"
)
