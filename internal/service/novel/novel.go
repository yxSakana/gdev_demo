package novel

import (
	novelDao "github.com/yxSakana/gdev_demo/internal/dao/novel"
	"github.com/yxSakana/gdev_demo/internal/model/do"
	"github.com/yxSakana/gdev_demo/internal/model/entity"
	"gorm.io/gorm"
)

func LinkNovelAndTags(db *gorm.DB, id uint64, tags []string) error {
	for _, tag := range tags {
		if tag == "" {
			continue
		}

		tagEntity, err := novelDao.GetNovelTagByNameWithAutoIncrement(db, tag)
		if err != nil {
			return err
		}
		ntrEntity := entity.NovelTagRel{
			NovelId:    id,
			NovelTagId: tagEntity.ID,
		}
		if err := novelDao.CreateNovelTagRel(db, &do.NovelTagRel{NovelTagRel: &ntrEntity}); err != nil {
			return err
		}
	}
	return nil
}
