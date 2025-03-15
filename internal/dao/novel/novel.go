package novel

import (
	"errors"

	"github.com/yxSakana/gdev_demo/internal/model/conv"
	"github.com/yxSakana/gdev_demo/internal/model/do"
	"github.com/yxSakana/gdev_demo/internal/model/entity"
	"gorm.io/gorm"
)

func CreateNovel(db *gorm.DB, novel *do.Novel) error {
	e := conv.NovelToEntity(*novel)
	err := db.Create(&e).Error
	if err == nil {
		novel.ID = e.ID
	}
	return err
}

func GetNovelByID(db *gorm.DB, id uint64) (novel *entity.Novel, err error) {
	novel = &entity.Novel{}
	err = db.First(novel, id).Error
	return
}

func CreateChapter(db *gorm.DB, chapter *do.NovelChapter) error {
	e := conv.NovelChapterToEntity(*chapter)
	err := db.Create(&e).Error
	if err == nil {
		chapter.ID = e.ID
	}
	return err
}

func CreateTag(db *gorm.DB, tag *entity.NovelTag) error {
	return db.Create(tag).Error
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
		err := CreateTag(db, &tag)
		if err != nil {
			return nil, err
		}
		return obj, nil
	}
	return obj, err
}

func CreateNovelTagRel(db *gorm.DB, ntr *entity.NovelTagRel) error {
	return db.Create(ntr).Error
}

func UpdateNovel(db *gorm.DB, nid uint64, updates map[string]any) error {
	return db.Model(&entity.Novel{}).Where("id = ?", nid).Updates(updates).Error
}

func DelNrt(db *gorm.DB, nid uint64) error {
	return db.Where("novel_id = ?", nid).Delete(&entity.NovelTagRel{}).Error
}
