package novel

import (
	"errors"
	"github.com/yxSakana/gdev_demo/internal/model/do"
	"github.com/yxSakana/gdev_demo/internal/model/entity"
	"gorm.io/gorm"
)

func CreateNovel(db *gorm.DB, novel *do.Novel) error {
	return db.Create(&novel.Novel).Error
}

func GetNovelByID(db *gorm.DB, id uint64) (novel *entity.Novel, err error) {
	novel = &entity.Novel{}
	err = db.First(novel, id).Error
	return
}

func CreateChapter(db *gorm.DB, chapter *do.NovelChapter) error {
	return db.Create(chapter.NovelChapter).Error
}

func CreateTag(db *gorm.DB, tag *do.NovelTag) error {
	return db.Create(tag.NovelTag).Error
}

func GetNovelTagByName(db *gorm.DB, name string) (*entity.NovelTag, error) {
	tag := &entity.NovelTag{}
	err := db.Where("name = ?", name).First(tag).Error
	if err != nil {
		return nil, err
	}
	return tag, nil
}

func GetNovelTagByNameWithAutoIncrement(db *gorm.DB, name string) (*entity.NovelTag, error) {
	obj, err := GetNovelTagByName(db, name)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		tag := entity.NovelTag{Name: name}
		obj := &tag
		err := CreateTag(db, &do.NovelTag{NovelTag: &tag})
		if err != nil {
			return nil, err
		}
		return obj, nil
	}
	return obj, err
}

func CreateNovelTagRel(db *gorm.DB, ntr *do.NovelTagRel) error {
	return db.Create(ntr.NovelTagRel).Error
}
