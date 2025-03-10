package image

import (
	"github.com/gin-gonic/gin"
	"github.com/yxSakana/gdev_demo/internal/dao"
	"gorm.io/gorm"
	"mime/multipart"

	imageDao "github.com/yxSakana/gdev_demo/internal/dao/image"
	"github.com/yxSakana/gdev_demo/internal/logic/user"
	"github.com/yxSakana/gdev_demo/internal/model"
	"github.com/yxSakana/gdev_demo/internal/model/do"
	"github.com/yxSakana/gdev_demo/internal/model/entity"
	service "github.com/yxSakana/gdev_demo/internal/service/image"
	"github.com/yxSakana/gdev_demo/utility"
)

func Create(c *gin.Context, in model.CreateImageCollectionInput) error {
	userEntity, err := user.GetUserinfo(c)
	if err != nil {
		return err
	}

	filePath, err := utility.GenerateFilePath(in.Cover)
	if err != nil {
		return err
	}

	if err := utility.CheckCoverFile(in.Cover, filePath); err != nil {
		return err
	}

	err = c.SaveUploadedFile(in.Cover, filePath)
	if err != nil {
		return err
	}

	collectionEntity := entity.ImageCollection{
		UserID:      userEntity.ID,
		Uploader:    userEntity.Username,
		Title:       in.Title,
		Description: in.Description,
		CoverUrl:    filePath,
	}

	db := dao.Ctx(c)
	d := &do.ImageCollection{ImageCollection: &collectionEntity}
	return db.Transaction(func(tx *gorm.DB) error {
		if err := imageDao.Create(tx, d); err != nil {
			return err
		}

		if err := service.LinkCollectionAndTags(tx, d.ID, in.Tags); err != nil {
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

	filePath, err := utility.GenerateFilePath(image)
	if err != nil {
		return err
	}

	if err := utility.CheckImageFile(image, filePath); err != nil {
		return err
	}

	if err := c.SaveUploadedFile(image, filePath); err != nil {
		return err
	}

	imgEntity := entity.Image{
		CollectionID: collectionID,
		ImageUrl:     filePath,
	}

	return db.Transaction(func(tx *gorm.DB) error {
		if err := imageDao.CreateImage(tx, &do.Image{Image: &imgEntity}); err != nil {
			return err
		}

		if err := imageDao.AutoUpdateCollectionImageCount(tx, collectionID); err != nil {
			return err
		}
		return nil
	})
}

func DetailImgCollectionByID(c *gin.Context, collectionID uint64) (*model.DetailImgCollectionOutput, error) {
	db := dao.Ctx(c)

	imgCollectionEntity, err := imageDao.GetCollectionByID(db, collectionID)
	if err != nil {
		return nil, err
	}

	tags, err := service.GetImgCollectionTags(db, collectionID)
	if err != nil {
		return nil, err
	}

	return &model.DetailImgCollectionOutput{
		ID:           collectionID,
		UploaderID:   imgCollectionEntity.UserID,
		UploaderName: imgCollectionEntity.Uploader,
		Title:        imgCollectionEntity.Title,
		Description:  imgCollectionEntity.Description,
		CoverUrl:     imgCollectionEntity.CoverUrl,
		Number:       imgCollectionEntity.Number,
		Tags:         tags,
		CreatedAt:    imgCollectionEntity.CreatedAt,
		UpdatedAt:    imgCollectionEntity.UpdatedAt,
	}, nil
}
