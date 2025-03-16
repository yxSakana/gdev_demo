package image

import (
	"gorm.io/gorm"

	imageDao "github.com/yxSakana/gdev_demo/internal/dao/image"
	"github.com/yxSakana/gdev_demo/internal/model/entity"
)

func LinkCollectionAndTags(db *gorm.DB, id uint64, tags []string) error {
	for _, tag := range tags {
		if tag == "" {
			continue
		}

		tagEntity, err := imageDao.GetTagByNameWithAutoIncrement(db, tag)
		if err != nil {
			return err
		}
		itrEntity := entity.ImageTagRel{
			CollectionID: id,
			ImageTagID:   tagEntity.ID,
		}
		if err := imageDao.CreateItr(db, &itrEntity); err != nil {
			return err
		}
	}
	return nil
}

func GetImgCollectionTags(db *gorm.DB, id uint64) ([]string, error) {
	var tags []string
	err := db.Model(&entity.ImageTag{}).
		Joins("LEFT JOIN image_tag_rel ON image_tags.id = image_tag_rel.image_tag_id").
		Joins("LEFT JOIN image_collections ON image_tag_rel.collection_id = image_collections.id").
		Where("image_collections.id = ?", id).
		Pluck("name", &tags).Error
	return tags, err
}
