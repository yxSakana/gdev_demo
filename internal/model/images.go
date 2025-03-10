package model

import (
	"mime/multipart"
	"time"
)

type CreateImageCollectionInput struct {
	Title       string                `form:"title" json:"title"`
	Description string                `form:"description" json:"description"`
	Cover       *multipart.FileHeader `form:"cover" json:"cover"`
	Tags        []string              `form:"tags" json:"tags"`
}

type DetailImgCollectionOutput struct {
	ID           uint64    `json:"id"`
	UploaderID   uint64    `json:"uploader_id"`
	UploaderName string    `json:"uploader_name"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	CoverUrl     string    `json:"cover_url"`
	Number       int       `json:"number"`
	Tags         []string  `json:"tags"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
