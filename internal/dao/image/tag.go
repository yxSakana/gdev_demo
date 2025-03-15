package image

import (
	"errors"
	"github.com/yxSakana/gdev_demo/internal/model/entity"
	"gorm.io/gorm"
)

func CreateTag(db *gorm.DB, imgTag *entity.ImageTag) error {
	return db.Create(imgTag).Error
}

func GetTagByName(db *gorm.DB, name string) (*entity.ImageTag, error) {
	imgTag := &entity.ImageTag{}
	if err := db.Where("name = ?", name).First(imgTag).Error; err != nil {
		return nil, err
	}
	return imgTag, nil
}

func GetTagByNameWithAutoIncrement(db *gorm.DB, name string) (*entity.ImageTag, error) {
	tagEntity, err := GetTagByName(db, name)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		tag := entity.ImageTag{Name: name}
		if err := CreateTag(db, &tag); err != nil {
			return nil, err
		}
		return &tag, nil
	}
	return tagEntity, err
}
