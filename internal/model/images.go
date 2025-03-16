package model

import (
	"mime/multipart"
)

type CreateImageCollectionInput struct {
	Title       string                `form:"title" json:"title"`
	Description string                `form:"description" json:"description"`
	Cover       *multipart.FileHeader `form:"cover" json:"cover"`
	Tags        []string              `form:"tags" json:"tags"`
}

type UpdateImageCollectionInput struct {
	Title       *string       `gorm:"title"`
	Description *string       `gorm:"description"`
	Cover       FileHeaderPtr `gorm:"-"`
	Tags        *[]string     `gorm:"-"`
}
