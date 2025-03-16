package consts

import (
	"errors"
	"gorm.io/gorm"
)

var ErrRecordNotFound = gorm.ErrRecordNotFound
var ErrCacheIsNil = errors.New("cache nil object") // 缓存nil对象以应对缓存穿透
