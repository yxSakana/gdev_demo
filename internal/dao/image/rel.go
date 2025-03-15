package image

import (
	"github.com/yxSakana/gdev_demo/internal/model/entity"
	"gorm.io/gorm"
)

func CreateItr(db *gorm.DB, itr *entity.ImageTagRel) error {
	return db.Create(itr).Error
}
