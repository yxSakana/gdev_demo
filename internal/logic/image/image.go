package image

import (
	"context"
	"errors"
	"log"
	"mime/multipart"
	"reflect"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/yxSakana/gdev_demo/internal/consts"
	"github.com/yxSakana/gdev_demo/internal/dao"
	imageDao "github.com/yxSakana/gdev_demo/internal/dao/image"
	"github.com/yxSakana/gdev_demo/internal/logic/user"
	"github.com/yxSakana/gdev_demo/internal/model"
	"github.com/yxSakana/gdev_demo/internal/model/conv"
	"github.com/yxSakana/gdev_demo/internal/model/do"
	"github.com/yxSakana/gdev_demo/internal/rediscon"
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

	icDo := new(do.ImageCollection)
	if err := icDo.DelCache(context.Background(), collectionID); err != nil {
		log.Printf("delete image failed, collectionID: %d, err: %v", collectionID, err)
	}

	return nil
}

func DetailImgCollectionByID(c *gin.Context, collectionID uint64) (icDo *do.ImageCollection, err error) {
	db := dao.Ctx(c)
	uid, err := user.GetUserID(c)
	if err != nil {
		return nil, err
	}
	icDo = new(do.ImageCollection)
	defer func() {
		if err == nil && icDo != nil {
			if err := rediscon.AddUvAndPv(context.Background(), uid, collectionID, icDo); err != nil {
				log.Printf("add Uv&Pv err: %v", err)
			}
			if err := imageDao.UpdateCollection(db, collectionID, map[string]any{"view": icDo.View}); err != nil {
				log.Printf("from cache update image: %#v", icDo)
			}
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := rediscon.NilCache(context.Background(), collectionID, icDo); err != nil {
				log.Printf("save to cache err: %v", err)
			}
		}
	}()

	if err := icDo.GetFromCache(context.Background(), collectionID); err == nil || errors.Is(err, consts.ErrCacheIsNil) {
		log.Printf("Get from cache image: %#v", icDo)
		return icDo, err
	}

	imgCollectionEntity, err := imageDao.GetCollectionByID(db, collectionID)
	if err != nil {
		return nil, err
	}

	tags, err := imgService.GetImgCollectionTags(db, collectionID)
	if err != nil {
		return nil, err
	}

	icDo = conv.ImageCollToDo(imgCollectionEntity, tags)
	if err := icDo.SaveToCache(context.Background(), collectionID); err != nil {
		log.Printf("save to cache err: %v", err)
	}
	return icDo, nil
}

func UpdateImageCollection(c *gin.Context, collectionID uint64, in model.UpdateImageCollectionInput) error {
	updateMap := make(map[string]interface{})

	t := reflect.TypeOf(in)
	v := reflect.ValueOf(in)
	for i := 0; i < t.NumField(); i++ {
		fieldVal := v.Field(i)
		if fieldVal.Kind() == reflect.Ptr && fieldVal.IsNil() {
			continue
		}
		field := t.Field(i)
		tag := field.Tag.Get("gorm")
		if tag == "-" {
			continue
		}

		updateMap[tag] = v.Field(i).Interface()
	}

	filePath, err := utility.SaveFile(c, in.Cover, utility.CoverFt)
	if err != nil && !errors.Is(err, utility.ErrFileHeaderIsNil) {
		log.Printf("save file err: %v", err)
		return err
	}
	updateMap["cover_url"] = filePath

	db := dao.Ctx(c)
	err = db.Transaction(func(tx *gorm.DB) error {
		if in.Tags != nil {
			if err := imageDao.DelItr(tx, collectionID); err != nil {
				return err
			}

			if err := imgService.LinkCollectionAndTags(tx, collectionID, *in.Tags); err != nil {
				return err
			}
		}

		return imageDao.UpdateCollection(tx, collectionID, updateMap)
	})
	if err != nil {
		return err
	}

	icDo := new(do.ImageCollection)
	if err := icDo.DelCache(context.Background(), collectionID); err != nil {
		log.Printf("del cache err: %v", err)
	}
	return nil
}
