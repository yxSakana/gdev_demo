package do

import (
	"github.com/yxSakana/gdev_demo/internal/model/entity"
)

type Novel struct {
	*entity.Novel
}

type NovelChapter struct {
	*entity.NovelChapter
}

type NovelTag struct {
	*entity.NovelTag
}

type NovelTagRel struct {
	*entity.NovelTagRel
}
