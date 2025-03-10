package image

import (
	"github.com/yxSakana/gdev_demo/internal/model/do"
	"gorm.io/gorm"
)

func CreateItr(db *gorm.DB, itr *do.ImageTagRel) error {
	return db.Create(&itr.ImageTagRel).Error
}
