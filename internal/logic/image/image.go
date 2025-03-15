package image

import (
	"errors"
	"github.com/yxSakana/gdev_demo/internal/consts"
	"github.com/yxSakana/gdev_demo/internal/rediscon"
	"log"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/yxSakana/gdev_demo/internal/dao"
	imageDao "github.com/yxSakana/gdev_demo/internal/dao/image"
	"github.com/yxSakana/gdev_demo/internal/logic/user"
	"github.com/yxSakana/gdev_demo/internal/model"
	"github.com/yxSakana/gdev_demo/internal/model/conv"
	"github.com/yxSakana/gdev_demo/internal/model/do"
	imgService "github.com/yxSakana/gdev_demo/internal/service/image"
	"github.com/yxSakana/gdev_demo/utility"
)

func Create(c *gin.Context, in model.CreateImageCollectionInput) error {
	userEntity, err := user.GetUserinfo(c)
	if err != nil {
		return err
	}

	filePath, err := utility.SaveFile(c, in.Cover, utility.CoverFt)

	collectionDo := do.ImageCollection{
		UserID:      userEntity.ID,
		Uploader:    userEntity.Username,
		Title:       in.Title,
		Description: in.Description,
		CoverUrl:    filePath,
	}

	db := dao.Ctx(c)
	return db.Transaction(func(tx *gorm.DB) error {
		if err := imageDao.Create(tx, &collectionDo); err != nil {
			return err
		}

		if err := imgService.LinkCollectionAndTags(tx, collectionDo.ID, in.Tags); err != nil {
			return err
		}
		return nil
	})
}

func UploadImage(c *gin.Context, collectionID uint64, image *multipart.FileHeader) error {
	db := dao.Ctx(c)

	if _, err := imageDao.GetCollectionByID(db, collectionID); err != nil {
		return err
	}

	filePath, err := utility.SaveFile(c, image, utility.ImageFt)

	imgDo := do.Image{
		CollectionID: collectionID,
		ImageUrl:     filePath,
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		if err := imageDao.CreateImage(tx, &imgDo); err != nil {
			return err
		}

		if err := imageDao.AutoUpdateCollectionImageCount(tx, collectionID); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	if err := imgService.DelCache(collectionID); err != nil {
		log.Printf("delete image failed, collectionID: %d, err: %v", collectionID, err)
	}

	return nil
}

func DetailImgCollectionByID(c *gin.Context, collectionID uint64) (*do.ImageCollection, error) {
	db := dao.Ctx(c)
	uid, err := user.GetUserID(c)
	if err != nil {
		return nil, err
	}

	cacheRet, err := imgService.GetFromCache(collectionID)
	if err == nil {
		log.Printf("cache ret is: %v", cacheRet)
		_ = rediscon.AddUvAndPv(uid, collectionID, consts.ImageCt)
		return cacheRet, nil
	}

	imgCollectionEntity, err := imageDao.GetCollectionByID(db, collectionID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			_ = imgService.SetCache(collectionID, nil)
		}
		return nil, err
	}

	tags, err := imgService.GetImgCollectionTags(db, collectionID)
	if err != nil {
		return nil, err
	}

	imgCollDo := conv.ImageCollToDo(imgCollectionEntity, tags)

	if err := imgService.SetCache(collectionID, imgCollDo); err != nil {
		log.Printf("ImageFt seriver SetCache err: %v", err)
	}

	_ = rediscon.AddUvAndPv(uid, collectionID, consts.ImageCt)
	return imgCollDo, nil
}
