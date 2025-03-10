package user

import (
	"github.com/yxSakana/gdev_demo/internal/model/do"
	"github.com/yxSakana/gdev_demo/internal/model/entity"
	"gorm.io/gorm"
)

func Create(db *gorm.DB, user *do.User) error {
	return db.Create(&user.User).Error
}

func GetUserByID(db *gorm.DB, id uint64) (result *entity.User, err error) {
	result = &entity.User{}
	err = db.First(result, id).Error
	return
}

func GetUserByUsername(db *gorm.DB, username string) (result *entity.User, err error) {
	result = &entity.User{}
	err = db.Where("username = ?", username).First(result).Error
	return
}

func Del() {

}
