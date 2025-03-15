package conv

import (
	"github.com/yxSakana/gdev_demo/internal/model/do"
	"github.com/yxSakana/gdev_demo/internal/model/entity"
)

func ImageCollToEntity(in *do.ImageCollection) *entity.ImageCollection {
	return &entity.ImageCollection{
		ID:          in.ID,
		UserID:      in.UserID,
		Uploader:    in.Uploader,
		Title:       in.Title,
		Description: in.Description,
		CoverUrl:    in.CoverUrl,
		Number:      in.Number,
		View:        in.View,
		Like:        in.Like,
		CreatedAt:   in.CreatedAt,
		UpdatedAt:   in.UpdatedAt,
	}
}

func ImageCollToDo(in *entity.ImageCollection, tags []string) *do.ImageCollection {
	return &do.ImageCollection{
		ID:          in.UserID,
		UserID:      in.UserID,
		Uploader:    in.Uploader,
		Title:       in.Title,
		Description: in.Description,
		CoverUrl:    in.CoverUrl,
		Tags:        tags,
		Number:      in.Number,
		View:        in.View,
		Like:        in.Like,
		CreatedAt:   in.CreatedAt,
		UpdatedAt:   in.UpdatedAt,
	}
}

func ImageToEntity(in *do.Image) *entity.Image {
	return &entity.Image{
		ID:           in.ID,
		CollectionID: in.CollectionID,
		ImageUrl:     in.ImageUrl,
	}
}
