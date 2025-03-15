package image

import (
	"github.com/yxSakana/gdev_demo/internal/model/conv"
	"github.com/yxSakana/gdev_demo/internal/model/do"
	"github.com/yxSakana/gdev_demo/internal/model/entity"
	"gorm.io/gorm"
)

func CreateImage(db *gorm.DB, img *do.Image) error {
	e := conv.ImageToEntity(img)
	err := db.Create(&e).Error
	if err == nil {
		img.ID = e.ID
	}
	return err
}

func GetCollectionImageCount(db *gorm.DB, id uint64) (count int64, err error) {
	err = db.Model(&entity.Image{}).
		Where("collection_id = ?", id).
		Count(&count).Error
	return
}
