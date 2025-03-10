package image

import (
	"github.com/yxSakana/gdev_demo/internal/model/do"
	"github.com/yxSakana/gdev_demo/internal/model/entity"
	"gorm.io/gorm"
)

func CreateImage(db *gorm.DB, img *do.Image) error {
	return db.Create(&img.Image).Error
}

func GetCollectionImageCount(db *gorm.DB, id uint64) (count int64, err error) {
	err = db.Model(&entity.Image{}).
		Where("collection_id = ?", id).
		Count(&count).Error
	return
}
