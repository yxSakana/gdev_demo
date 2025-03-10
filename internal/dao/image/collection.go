package image

import (
	"github.com/yxSakana/gdev_demo/internal/model/do"
	"github.com/yxSakana/gdev_demo/internal/model/entity"
	"gorm.io/gorm"
)

func Create(db *gorm.DB, imgColl *do.ImageCollection) error {
	return db.Create(&imgColl.ImageCollection).Error
}

func GetCollectionByID(db *gorm.DB, id uint64) (*entity.ImageCollection, error) {
	ret := &entity.ImageCollection{}
	err := db.First(ret, id).Error
	return ret, err
}

func UpdateCollection(db gorm.DB, id uint64, fields map[string]any) error {
	return db.Model(&entity.ImageCollection{ID: id}).Updates(fields).Error
}

func AutoUpdateCollectionImageCount(db *gorm.DB, id uint64) error {
	return db.Exec(`
	UPDATE image_collections SET number = (
		SELECT COUNT(*) FROM images WHERE collection_id = ?)
	WHERE image_collections.id = ?`, id, id).Error
}
