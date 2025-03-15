package v1

import "mime/multipart"

// 获取某个用户的所有n ----> user

type FileHeaderPtr = *multipart.FileHeader

type CreateNovelReq struct {
	Title       string        `form:"title" binding:"required" json:"title"`
	Description string        `form:"description" json:"description"`
	Cover       FileHeaderPtr `form:"cover" json:"cover_url"`
	Tags        []string      `form:"tags" json:"tags"`
}
type CreateNovelRes struct {
	NovelID uint64 `json:"novel_id"`
}

type UploadChapterReq struct {
	NovelId uint64 `form:"novel_id" binding:"required" json:"novel_id"`
	Title   string `form:"title" binding:"required" json:"title"`
	Number  int    `form:"number" binding:"required" json:"number"`
	Content string `form:"content" binding:"required" json:"content"`
}
type UploadChapterRes struct {
}

type DetailNovelReq struct {
	NovelID uint64 `form:"novel_id" json:"novel_id"`
}
type DetailNovelRes struct {
	NovelDetail
}

type QueryReq struct {
	Id        *uint64 `json:"id" form:"id"`
	Tag       *string `json:"tag" form:"tag"`
	Author    *string `json:"author" form:"author"`
	Title     *string `json:"title" form:"title"`
	WordCount *uint   `json:"word_count" form:"word_count"`
	View      *uint   `json:"view" form:"view"`
	Like      *uint   `json:"like" form:"like"`
}
type QueryRes struct {
	Total uint          `json:"total"`
	List  []NovelDetail `json:"list"`
}

type UpdateNovelReq struct {
	Title       *string       `form:"title" gorm:"title"`
	Description *string       `form:"description" gorm:"description"`
	Cover       FileHeaderPtr `form:"cover" gorm:"-"`
	Status      *int          `form:"status" gorm:"status"`
	Tags        *[]string     `form:"tags" gorm:"-"`
}
type UpdateNovelRes struct{}
