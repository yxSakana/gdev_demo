package v1

import "mime/multipart"

type FileHeaderPtr = *multipart.FileHeader

type CreateReq struct {
	Title       string        `form:"title" binding:"required" json:"title"`
	Description string        `form:"description" json:"description"`
	Cover       FileHeaderPtr `form:"cover" json:"cover"`
	Tags        []string      `form:"tags" json:"tags"`
}
type CreateRes struct {
}

type UploadImageReq struct {
	CollectionID uint64        `form:"collection_id" binding:"required" json:"collection_id"`
	Image        FileHeaderPtr `form:"image" binding:"required" json:"image"`
}
type UploadImageRes struct {
}

type UploadImagesReq struct {
	CollectionID uint64          `form:"collection_id" binding:"required" json:"collection_id"`
	Images       []FileHeaderPtr `form:"images" binding:"required" json:"images"`
}
type UploadImagesRes struct{}

type DetailImgCollectionReq struct {
	CollectionID uint64 `form:"collection_id" binding:"required" json:"collection_id"`
}
type DetailImgCollectionRes struct {
	ImageCollection
}

type UpdateImageCollectionReq struct {
	Title       *string       `form:"title" binding:"required" json:"title"`
	Description *string       `form:"description" binding:"required" json:"description"`
	Cover       FileHeaderPtr `form:"cover" binding:"required" json:"cover"`
	Tags        *[]string     `form:"tags" binding:"required" json:"tags"`
}
type UpdateImageCollectionRes struct {
}

type DelImageCollectionReq struct {
	Id uint64 `json:"id" binding:"required"`
}
type DelImageCollectionRes struct{}
