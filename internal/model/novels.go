package model

import (
	"mime/multipart"
)

type FileHeaderPtr = *multipart.FileHeader

type NovelQueryInput struct {
	Id        *uint64 `json:"id" form:"id"`
	Tag       *string `json:"tag" form:"tag"`
	Author    *string `json:"author" form:"author"`
	Title     *string `json:"title" form:"title"`
	WordCount *uint   `json:"word_count" form:"word_count"`
	View      *uint   `json:"view" form:"view"`
	Like      *uint   `json:"like" form:"like"`
}

type UpdateNovelInput struct {
	Title         *string       `gorm:"title"`
	Description   *string       `gorm:"description"`
	Cover         FileHeaderPtr `gorm:"-"`
	Status        *int          `gorm:"status"`
	Tags          *[]string     `gorm:"-"`
	ChapterNumber *uint         `gorm:"chapter_number"`
	WordCount     *uint         `gorm:"word_count"`
	View          *uint         `gorm:"view"`
	Like          *uint         `gorm:"like"`
}
